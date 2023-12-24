package org.acme;

import jakarta.ws.rs.GET;
import jakarta.ws.rs.Path;
import jakarta.ws.rs.PathParam;
import jakarta.ws.rs.Produces;
import jakarta.ws.rs.core.MediaType;

@Path("/")
public class GreetingResource {

    @GET
    @Path("/{path:.*}")
    @Produces(MediaType.TEXT_PLAIN)
    public String helloFromPath(@PathParam("path") String path) {
        if(path == ""){
            return "try to add /test as path\n";
        } else {
            return "hello from /" + path;
        }
    }
}
