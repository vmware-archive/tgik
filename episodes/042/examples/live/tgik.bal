import ballerina/http;
import ballerina/log;
import ballerinax/kubernetes;




@kubernetes:Service {
    serviceType: "NodePort",
    name: "tgik-ballerina"
}
endpoint http:Listener listener {
    port: 9090
};

@kubernetes:Deployment {
    apiVersion: "v1/deployment",
    enableLiveness:true,
    image: "krisnova/tgik-ballerina",
    name: "tgik-ballerina",
    imagePullPolicy: "Always"
}
@http:ServiceConfig {
    basePath: "/"
}
service<http:Service> hello bind listener {

    sayHello(endpoint caller, http:Request req) {

        http:Response res = new;

        res.setPayload("Hello, World!");

        caller->respond(res) but { error e => log:printError(
                                                  "Error sending response", err = e) };
    }

    @http:ResourceConfig {
        methods: ["POST"],
        path: "/sayGoodbye"
    }
    sayGoodbye(endpoint caller, http:Request req) {

        http:Response res = new;

        res.setPayload("Goodbye, World!");

        caller->respond(res) but { error e => log:printError(
                                                  "Error sending response", err = e) };
    }

}