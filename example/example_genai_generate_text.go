// Copyright (c) 2016, 2018, 2024, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.

// Example code for using the Gen AI Generate Text API

package example

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/example/helpers"
	"github.com/oracle/oci-go-sdk/v65/generativeaiinference"
)

func ExampleGenerateText() {
	region := "us-chicago-1"
	compartment_id := "" // compartmentId that has policies grant permissions for using Generative AI Service
	c, clerr := generativeaiinference.NewGenerativeAiInferenceClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	c.SetRegion(region)
	ctx := context.Background()

	modelId := "cohere.command"
	prompt := "Tell me a fact about the Earth"
	max_tokens := 75
	temperature := 0.75
	frequency_penalty := 1.0
	top_p := 0.7
	isStream := false
	details := generativeaiinference.GenerateTextDetails{
		CompartmentId: &compartment_id,
		ServingMode: generativeaiinference.OnDemandServingMode{
			ModelId: &modelId,
		},
		InferenceRequest: generativeaiinference.CohereLlmInferenceRequest{
			Prompt:           &prompt,
			IsStream:         &isStream,
			MaxTokens:        &max_tokens,
			Temperature:      &temperature,
			TopP:             &top_p,
			FrequencyPenalty: &frequency_penalty,
		},
	}

	req := generativeaiinference.GenerateTextRequest{
		GenerateTextDetails: details,
	}

	// application/json response type
	resp, err := c.GenerateText(ctx, req)
	helpers.FatalIfError(err)
	fmt.Println(resp) // Response body contains generated text response to the prompt

	// text/event-stream response type
	isStream = true
	resp, err = c.GenerateText(ctx, req)
	helpers.FatalIfError(err)

	sseReader, err := common.NewSSEReader(resp.HTTPResponse())
	helpers.FatalIfError(err)

	// Print each event's text field
	err = sseReader.ReadAllEvents(func(event []byte) {
		var parsed generativeaiinference.GeneratedText
		json.Unmarshal([]byte(event), &parsed)
		fmt.Print(*parsed.Text)
	})
	helpers.FatalIfError(err)

}
