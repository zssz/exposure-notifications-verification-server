// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Exchanges a verification token for a verification certificate (step 2).
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/google/exposure-notifications-verification-server/pkg/clients"
)

func main() {
	tokenFlag := flag.String("token", "", "verification token to claim")
	hmacFlag := flag.String("hmac", "", "data to certify")
	apikeyFlag := flag.String("apikey", "", "API Key to use")
	addrFlag := flag.String("addr", "http://localhost:8080", "protocol, address and port on which to make the API call")
	timeoutFlag := flag.Duration("timeout", 5*time.Second, "request time out duration in the format: 0h0m0s")
	flag.Parse()

	ctx := context.Background()

	request, response, err := clients.GetCertificate(ctx, *addrFlag, *apikeyFlag, *tokenFlag, *hmacFlag, *timeoutFlag)
	if err != nil {
		log.Printf("Send request: %+v", request)
		log.Fatalf("error making API call: %v", err)
	}
	log.Printf("Result: \n%+v", response)
}
