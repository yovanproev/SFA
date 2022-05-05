/*
 *
 * Copyright 2015 gRPC authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/examples/helloworld/greeter_server/db"
	"google.golang.org/grpc/examples/helloworld/greeter_server/hns"

	pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

var (
	port = flag.Int("port", 9000, "The server port")
)

type server struct {
	pb.UnimplementedTopStoriesServiceServer
}

func (s *server) FetchStories(context.Context, *pb.Story) (*pb.My_Replies, error) {
	sqLite, resultFromDb := hns.InitializeDB("../data.db")

	var topTenStories hns.SummedType
	var err error

	if resultFromDb.Items == nil {
		topTenStories, err = hns.WriteToDBAndPush(context.Background(), sqLite)
		if err != nil {
			log.Println(err)
		}
	} else {
		q := db.New(sqLite)
		lastStoredItems, err := q.GetLastStory(context.Background())
		if err != nil {
			log.Println(err)
		}
		now := time.Now()
		hourlyCheck := lastStoredItems.DateStamp.Add(1 * time.Hour)

		if now.After(hourlyCheck) {
			topTenStories, err = hns.WriteToDBAndPush(context.Background(), sqLite)
			if err != nil {
				log.Println(err)
			}
		} else {
			topTenStories = resultFromDb
		}
	}

	var r []*pb.Reply
	for _, v := range topTenStories.Items {
		rep := pb.Reply{
			Title: v.Title,
			Score: v.Score,
		}
		r = append(r, &rep)
	}

	sliceOfReplies := pb.My_Replies{
		Reply: r,
	}

	return &sliceOfReplies, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTopStoriesServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

// 2022/05/05 17:09:52 Top stories: [
//Title:"Mermaid: Create diagrams and visualizations using text and code"  Score:8
//Title:"Modern Python Performance Considerations"  Score:58
//Title:"Mechanical Watch"  Score:3040
//Title:"USB Cheat Sheet"  Score:186
//Title:"Teen mental health is plummeting and social media is a major contributing cause [pdf]"  Score:578
//Title:"Shopify to Acquire Delivrr for $2.1B"  Score:73
// Title:"How Crossrail was affected by the curvature of the Earth (2018)"  Score:67
//Title:"I Accidentally Deleted 7TB of Videos Before Going to Production"  Score:263
//Title:"JavaScript Containers"  Score:80
//Title:"Apple, Google and Microsoft Commit to Expanded Support for FIDO Standard"  Score:82
//]
