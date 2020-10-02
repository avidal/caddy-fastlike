use std::convert::TryFrom;

use fastly::http::{Method, StatusCode};
use fastly::{Body, Error, Request, RequestExt, Response, ResponseExt};

const BACKEND: &str = "backend";

#[fastly::main]
fn main(req: Request<Body>) -> Result<impl ResponseExt, Error> {
    match (req.method(), req.uri().path()) {
        (&Method::GET, "/teapot") => Ok(Response::builder()
            .status(StatusCode::IM_A_TEAPOT)
            .body(Body::try_from("Hello, teapot!")?)?
        ),

        _ => req.send(BACKEND)
    }
}
