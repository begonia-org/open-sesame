package main

import "fmt"

func main() {
	fmt.Printf("api.Version")
	// 连接到 gRPC 服务器
	// conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalf("did not connect: %v", err)
	// }
	// defer conn.Close()

	// // 创建 gRPC 客户端
	// client := api.NewDocumentEmbedServiceClient(conn)
	// stream, err := client.Embed(context.Background(), &api.DocumentEmbedRequest{
	// 	FileId:    "458092175117258752",
	// 	Lang:      "Chinese",
	// 	ChunkType: api.ChunkType_CHUNK_TYPE_NAIVE,
	// 	LlmType:   api.LLMType_LLM_TYPE_EMBEDDING,
	// 	ParserConfig: &api.ParserConfig{
	// 		FilenameEmbdWeight: 0.1,
	// 		TaskPageSize:       12,
	// 		DoLayoutRecognize:  true,
	// 	},
	// 	Llm: &api.LLM{
	// 		Factory:     "BAAI",
	// 		ModelType:   api.LLMType_LLM_TYPE_EMBEDDING.String(),
	// 		Name:        "BAAI/bge-large-zh-v1.5",
	// 		LimitTokens: 1024 * 1024,
	// 	},
	// })
	// if err != nil {
	// 	log.Fatalf("embed error:%v", err)
	// }
	// for {
	// 	resp, err := stream.Recv()
	// 	if err != nil {
	// 		log.Fatalf("embed error:%v", err)
	// 		break
	// 	}
	// 	log.Printf("embed response:%v", resp)
	// }
	// md, ok := metadata.FromIncomingContext(stream.Context())
	// if ok {
	// 	log.Printf("metadata:%v", md)
	// }

}
