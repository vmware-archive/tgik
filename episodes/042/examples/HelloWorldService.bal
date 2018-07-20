import ballerina/http;
import ballerina/log;
import ballerinax/kubernetes;



@kubernetes:Service {
    serviceType: "NodePort",
    name: "ballerina-demo"
}
endpoint http:Listener listener {
    port: 9090
};


@kubernetes:Deployment {
    image: "krisnova/ballerina-demo",
    name: "ballerina-demo"
}

@http:ServiceConfig {
    basePath: "/"
}
service<http:Service> hello bind { port: 9090 } {

    sayHello(endpoint caller, http:Request req) {
        http:Response res = new;

        res.setPayload("Hello, World!");

        caller->respond(res) but { error e => log:printError(
                           "Error sending response", err = e) };
    }
}