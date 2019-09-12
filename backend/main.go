package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"

	pb "github.com/victorsteven/grpc-speech/api"
)

type server struct{}

func main() {

	port := flag.Int("p", 8080, "Port to listen to")
	flag.Parse()

	logrus.Infof("listening to port %d", *port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()

	pb.RegisterTextToSpeechServer(s, server{})

	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("could not serve: %v", err)
	}
	// cmd := exec.Command("flite", "-t", os.Args[1], "-o", "output.wav")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {

	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("could not create tmp file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("could not close %s: %v", f.Name(), err)
	}
	// return nil, fmt.Errorf("not implemented")
	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())

	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("could not read tmp file: %v", err)
	}
	return &pb.Speech{Audio: data}, nil
}
