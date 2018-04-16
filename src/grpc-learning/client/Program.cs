using System;
using Grpc;
using Grpc.Core;
using static Grpc.grpc;

namespace client {
    class Program {
        static void Main (string[] args) {
            Console.WriteLine ("grpc client:Hello World!");
            Channel chan = new Channel ("127.0.0.1:9090", ChannelCredentials.Insecure);
            grpcClient client = new grpcClient (chan);
            var request = new HelloRequest () {
                Name = "wuxian"
            };
            while (true) {
                var stdinStr = Console.ReadLine ();
                request.Name = stdinStr;
                var response = client.SayHello (request);
                System.Console.WriteLine("response:"+response.Message);
            }
        }
    }
}