package ycsdk

import (
	"github.com/yandex-cloud/go-sdk/gen/apigateway"
	"github.com/yandex-cloud/go-sdk/gen/apigateway/websocket"
	"github.com/yandex-cloud/go-sdk/gen/containers"
	"github.com/yandex-cloud/go-sdk/gen/eventrouter"
	"github.com/yandex-cloud/go-sdk/gen/functions"
	"github.com/yandex-cloud/go-sdk/gen/mdbproxy"
	"github.com/yandex-cloud/go-sdk/gen/triggers"
	"github.com/yandex-cloud/go-sdk/gen/workflows"
)

type Serverless struct {
	sdk *SDK
}

const (
	FunctionServiceID             Endpoint = "serverless-functions"
	TriggerServiceID              Endpoint = "serverless-triggers"
	APIGatewayServiceID           Endpoint = "serverless-apigateway"
	MDBProxyServiceID             Endpoint = "mdbproxy"
	ServerlessContainersServiceID Endpoint = "serverless-containers"
	APIGatewayWebsocketServiceID  Endpoint = "apigateway-connections"
	EventrouterServiceID          Endpoint = "serverless-eventrouter"
	EventrouterEventsServiceID    Endpoint = "serverlesseventrouter-events"
	WorkflowServiceID             Endpoint = "serverless-workflows"
)

func (s *Serverless) Functions() *functions.Function {
	return functions.NewFunction(s.sdk.getConn(FunctionServiceID))
}

func (s *Serverless) Triggers() *triggers.Trigger {
	return triggers.NewTrigger(s.sdk.getConn(TriggerServiceID))
}

func (s *Serverless) APIGateway() *apigateway.Apigateway {
	return apigateway.NewApigateway(s.sdk.getConn(APIGatewayServiceID))
}

func (s *Serverless) MDBProxy() *mdbproxy.Proxy {
	return mdbproxy.NewProxy(s.sdk.getConn(MDBProxyServiceID))
}

func (s *Serverless) Containers() *containers.Container {
	return containers.NewContainer(s.sdk.getConn(ServerlessContainersServiceID))
}

func (s *Serverless) APIGatewayWebsocket() *websocket.Websocket {
	return websocket.NewWebsocket(s.sdk.getConn(APIGatewayWebsocketServiceID))
}

func (s *Serverless) Eventrouter() *eventrouter.Eventrouter {
	return eventrouter.NewEventrouter(s.sdk.getConn(EventrouterServiceID))
}

func (s *Serverless) EventrouterEvents() *eventrouter.Eventrouter {
	return eventrouter.NewEventrouter(s.sdk.getConn(EventrouterEventsServiceID))
}

func (s *Serverless) Workflow() *workflows.Workflow {
	return workflows.NewWorkflow(s.sdk.getConn(WorkflowServiceID))
}
