package main

import (
        "fmt"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/SkyAPM/go2sky"
	httpPlugin "github.com/SkyAPM/go2sky/plugins/http"
	"github.com/SkyAPM/go2sky/reporter"
)

var (
	grpc        bool
	oapServer   string
	upstreamURL string
	listenAddr  string
	serviceName string

	client *http.Client
)

func init() {
	flag.BoolVar(&grpc, "grpc", true, "use grpc reporter")
	flag.StringVar(&oapServer, "oap-server", "192.168.10.244:11800", "oap server address")
	flag.StringVar(&upstreamURL, "http://192.168.10.234:38900/massrelayapi/actuator/prometheus", "upstream-service", "upstream service url")
	flag.StringVar(&listenAddr, "listen-addr", ":8080", "listen address")
	flag.StringVar(&serviceName, "service-test", "go2sky", "service name")
}

func ServerHTTP(writer http.ResponseWriter, request *http.Request) {
	time.Sleep(time.Duration(500) * time.Millisecond)
	go2sky.PutCorrelation(request.Context(), "MIDDLE_KEY", "go2sky")

	clientReq, err := http.NewRequest(http.MethodPost, upstreamURL, nil)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("unable to create http request error: %v \n", err)
		return
	}
	clientReq = clientReq.WithContext(request.Context())
	res, err := client.Do(clientReq)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("unable to do http request error: %v \n", err)
		return
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		log.Printf("read http response error: %v \n", err)
		return
	}
	writer.WriteHeader(res.StatusCode)
	_, _ = writer.Write(body)
}

func main() {
	flag.Parse()
	fmt.Println("test1")
	var report go2sky.Reporter
	var err error
	if grpc {
		// report, err = reporter.NewGRPCReporter(oapServer)
		report, err = reporter.NewGRPCReporter("192.168.10.244:11800")
		if err != nil {
			log.Fatalf("crate grpc reporter error: %v \n", err)
		}
	} else {
		report, err = reporter.NewLogReporter()
		if err != nil {
			log.Fatalf("crate log reporter error: %v \n", err)
		}
	}
	fmt.Println("test2")

	tracer, err := go2sky.NewTracer(serviceName, go2sky.WithReporter(report))
	if err != nil {
		log.Fatalf("crate tracer error: %v \n", err)
	}

	client, err = httpPlugin.NewClient(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}

	route := http.NewServeMux()
	route.HandleFunc("/", ServerHTTP)

	fmt.Println("test3")
	sm, err := httpPlugin.NewServerMiddleware(tracer)
	if err != nil {
		log.Fatalf("create client error %v \n", err)
	}
	err = http.ListenAndServe(listenAddr, sm(route))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test4")
}

