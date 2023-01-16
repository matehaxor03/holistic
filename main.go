package main

import (
	webserver "github.com/matehaxor03/holistic_webserver/webserver"
	processor "github.com/matehaxor03/holistic_processor/processor"
	queue "github.com/matehaxor03/holistic_queue/queue"
	common "github.com/matehaxor03/holistic_common/common"
	"fmt"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var errors []error
	queue_server, queue_server_errors := queue.NewQueueServer("5000", "server.crt", "server.key", "127.0.0.1", "5002")
	if queue_server_errors != nil {
		errors = append(errors, queue_server_errors...)
	} else if common.IsNil(queue_server) {
		errors = append(errors, fmt.Errorf("queue is nil"))
	}

	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}

	complete_function := queue_server.GetCompleteFunction()
	get_next_message_function := queue_server.GetNextMessageFunction()
	push_message_function := queue_server.GetPushBackFunction()


	web_server, web_server_errors := webserver.NewWebServer(push_message_function, "5001", "server.crt", "server.key", "127.0.0.1", "5000")
	if web_server_errors != nil {
		errors = append(errors, web_server_errors...)	
	} else if common.IsNil(web_server) {
		errors = append(errors, fmt.Errorf("web_server is nil"))
	} 

	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}
	
	processor_server, processor_server_errors := processor.NewProcessorServer(complete_function, get_next_message_function, push_message_function, "5002", "server.crt", "server.key", "127.0.0.1", "5000")
	if processor_server_errors != nil {
		errors = append(errors, processor_server_errors...)	
	} else if common.IsNil(processor_server) {
		errors = append(errors, fmt.Errorf("web_server is nil"))
	} 

	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}

	processor_wakeup_function := processor_server.GetWakeupProcessorFunction()
	queue_server.SetProcessorCallbackFunction(processor_wakeup_function)

	go web_server.Start()
	go queue_server.Start()
	go processor_server.Start()
	

	wg.Add(1)
	wg.Wait()
	os.Exit(0)
}