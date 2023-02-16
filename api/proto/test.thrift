namespace go tiktok

struct TestRequest {
  1: string name
}

struct TestResponse {
  1: string message
}

service TestService {
  TestResponse Test(1: TestRequest request)
}
