// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package zedcloudapi

import (
    "fmt"
	"github.com/go-test/deep"
	"github.com/zededa/zedcloud-api/swagger_models"
	"testing"
)

func expectedOutput(rsp *swagger_models.ZsrvResponse) error {
	errStr := ""
    if rsp.HTTPStatusCode != 200 {
        errStr = fmt.Sprintf("Request Failed. HTTPStatusCode: %d, HTTPStatusMsg: %s",
            rsp.HTTPStatusCode, rsp.HTTPStatusMsg)
    }
	for _, zerr := range rsp.Error {
		errCode := ""
		if zerr.Ec != nil  && *zerr.Ec != swagger_models.ZsrvErrorCodeZMsgSucess {
			errCode = string(*zerr.Ec)
            errStr += " " + fmt.Sprintf("ErrorCode: %s, ErrorDetails: %s\n",
                errCode, zerr.Details)
		}
	}
    if errStr != "" {
        return fmt.Errorf("%s", errStr)
    }
    return nil
}

var zsrvNoErrEntries *swagger_models.ZsrvResponse = &swagger_models.ZsrvResponse{
    EndTime: "Test End Time",
    HTTPStatusCode: 200,
    HTTPStatusMsg: "SUCCESS",
    Error: []*swagger_models.ZsrvError{},
}
var zsrvOneSuccessEntry *swagger_models.ZsrvResponse = &swagger_models.ZsrvResponse{
    EndTime: "Test End Time",
    HTTPStatusCode: 200,
    HTTPStatusMsg: "SUCCESS",
    Error: []*swagger_models.ZsrvError{
        &swagger_models.ZsrvError{
            Ec: swagger_models.NewZsrvErrorCode(
                swagger_models.ZsrvErrorCodeZMsgSucess),
            Details: "",
        },
    },
}

var zsrvOneErrorEntry *swagger_models.ZsrvResponse = &swagger_models.ZsrvResponse{
    EndTime: "Test End Time",
    HTTPStatusCode: 409,
    HTTPStatusMsg: "Conflict",
    Error: []*swagger_models.ZsrvError{
        &swagger_models.ZsrvError{
            Ec: swagger_models.NewZsrvErrorCode(
                swagger_models.ZsrvErrorCodeAlreadyExists),
            Details: "Object Already Exists",
        },
    },
}

var zsrvMultipleErrorEntries *swagger_models.ZsrvResponse = &swagger_models.ZsrvResponse{
    EndTime: "Test End Time",
    HTTPStatusCode: 409,
    HTTPStatusMsg: "Conflict",
    Error: []*swagger_models.ZsrvError{
        &swagger_models.ZsrvError{
            Ec: swagger_models.NewZsrvErrorCode(
                swagger_models.ZsrvErrorCodeAlreadyExists),
            Details: "Object Already Exists",
        },
        &swagger_models.ZsrvError{
            Ec: swagger_models.NewZsrvErrorCode(
                swagger_models.ZsrvErrorCodeIncompleteData),
            Details: "IncompleteData",
        },
    },
}

func TestProcessZsrvResponseError(t *testing.T) {
	cases := []struct {
		description             string
		input                   interface{}
        expectedOutput          error
	}{
		{
			description:             "NON-ZsrvResponse",
			input:                   &swagger_models.NetInstConfig{},
            expectedOutput: nil,
		},
		{
			description:             "Empty ZsrvResponse",
			input:                   &swagger_models.ZsrvResponse{},
            expectedOutput: expectedOutput(&swagger_models.ZsrvResponse{}),
		},
		{
			description:             "ZsrvResponse with No Error entries",
			input:                   &swagger_models.ZsrvResponse{
                EndTime: "Test End Time",
                HTTPStatusCode: 200,
                HTTPStatusMsg: "SUCCESS",
            },
            expectedOutput: nil,
		},
		{
			description:             "Empty Error List",
			input:                   zsrvNoErrEntries,
            expectedOutput: expectedOutput(zsrvNoErrEntries),
		},
		{
			description:             "Error List with ONE Success Entry",
			input:                   zsrvOneSuccessEntry,
            expectedOutput: expectedOutput(zsrvOneSuccessEntry),
		},
		{
			description:             "Error List with one Error Entry",
			input:                   zsrvOneErrorEntry,
            expectedOutput: expectedOutput(zsrvOneErrorEntry),
		},
		{
			description:             "Error List with Multiple Error Entries",
			input:                   zsrvMultipleErrorEntries,
            expectedOutput: expectedOutput(zsrvMultipleErrorEntries),
		},
	}

	for _, c := range cases {
        opErr := processZsrvResponseError(c.input)
		if diff := deep.Equal(c.expectedOutput, opErr); diff != nil {
                t.Fatalf("Test Failed: %s\nExpected: %+v\n Actual: %+v\n" +
                    "diff: %+v", c.description, c.expectedOutput, opErr, diff)
        }
	}
}
