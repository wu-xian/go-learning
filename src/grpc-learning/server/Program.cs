using System;
using System.Threading.Tasks;
using Grpc.Core;

namespace server
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("grpc server:Hello World!");
            Server server = new Server(){
                Services = {Grpc.grpc.BindService(new gRPCImpl())},
                Ports = {new ServerPort("localhost",9090,ServerCredentials.Insecure)}
            };
            server.Start();
            Console.WriteLine("gRPC server listening on port 9090");
            Console.WriteLine("任意键退出...");
            Console.ReadKey();

            server.ShutdownAsync().Wait();
        }
    }

        class gRPCImpl : Grpc.grpc.grpcBase
    {
        // 实现SayHello方法
        public override Task<Grpc.HelloResponse> SayHello(Grpc.HelloRequest request, ServerCallContext context)
        {
            System.Console.WriteLine("Echo message:"+request.Name);
            return Task.FromResult(new Grpc.HelloResponse { Message = "Hello " + request.Name });
        }
    }
}
