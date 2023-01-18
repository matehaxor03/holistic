package main

import (
	webserver "github.com/matehaxor03/holistic_webserver/webserver"
	processor "github.com/matehaxor03/holistic_processor/processor"
	json "github.com/matehaxor03/holistic_json/json"
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

	queue_complete_functions := make(map[string](*func(json.Map) []error))
	queue_get_next_message_functions := make(map[string]((*func(string) (json.Map, []error))))
	queue_push_back_functions :=  make(map[string](*func(json.Map) (*json.Map,[]error)))

	queue_controllers := queue_server.GetControllers()
	for queue_controller_name, queue_controller := range queue_controllers {
		queue_complete_functions[queue_controller_name] = queue_controller.GetCompleteFunction()
		queue_get_next_message_functions[queue_controller_name] = queue_controller.GetNextMessageFunction()
		queue_push_back_functions[queue_controller_name] = queue_controller.GetPushBackFunction()
	}

	web_server, web_server_errors := webserver.NewWebServer("5001", "server.crt", "server.key", "127.0.0.1", "5000")
	if web_server_errors != nil {
		errors = append(errors, web_server_errors...)	
	} else if common.IsNil(web_server) {
		errors = append(errors, fmt.Errorf("web_server is nil"))
	} 

	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}
	
	processor_server, processor_server_errors := processor.NewProcessorServer("5002", "server.crt", "server.key", "127.0.0.1", "5000")
	if processor_server_errors != nil {
		errors = append(errors, processor_server_errors...)	
	} else if common.IsNil(processor_server) {
		errors = append(errors, fmt.Errorf("web_server is nil"))
	} else {
		processor_server.SetQueueCompleteFunctions(queue_complete_functions)
		processor_server.SetQueuePushBackFunctions(queue_push_back_functions)
		processor_server.SetQueueGetNextMessageFunctions(queue_get_next_message_functions)
	}

	if len(errors) > 0 {
		fmt.Println(errors)
		os.Exit(1)
	}

	processor_manager_wakeup_functions := make(map[string](*func()))
	processor_controllers := processor_server.GetControllers()
	for processor_controller_name, processor_controller := range processor_controllers {
		processor_manager_wakeup_functions[processor_controller_name] = processor_controller.GetWakeupProcessorManagerFunction()
	}
	queue_server.SetProcessorWakeUpFunctions(processor_manager_wakeup_functions)

	go web_server.Start()
	go queue_server.Start()
	go processor_server.Start()
	

	wg.Add(1)
	wg.Wait()
	os.Exit(0)
}