// Code generated by smithy-go-codegen DO NOT EDIT.

package query

import (
	"bytes"
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/internal/protocoltest/query/types"
	"github.com/awslabs/smithy-go/middleware"
	"github.com/awslabs/smithy-go/ptr"
	smithyrand "github.com/awslabs/smithy-go/rand"
	smithytesting "github.com/awslabs/smithy-go/testing"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestClient_QueryMaps_awsAwsquerySerialize(t *testing.T) {
	cases := map[string]struct {
		Params        *QueryMapsInput
		ExpectMethod  string
		ExpectURIPath string
		ExpectQuery   []smithytesting.QueryItem
		RequireQuery  []string
		ForbidQuery   []string
		ExpectHeader  http.Header
		RequireHeader []string
		ForbidHeader  []string
		BodyMediaType string
		BodyAssert    func(io.Reader) error
	}{
		// Serializes query maps
		"QuerySimpleQueryMaps": {
			Params: &QueryMapsInput{
				MapArg: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&MapArg.entry.1.key=bar
			&MapArg.entry.1.value=Bar
			&MapArg.entry.2.key=foo
			&MapArg.entry.2.value=Foo`))
			},
		},
		// Serializes query maps and uses xmlName
		"QuerySimpleQueryMapsWithXmlName": {
			Params: &QueryMapsInput{
				RenamedMapArg: map[string]string{
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&Foo.entry.1.key=foo
			&Foo.entry.1.value=Foo`))
			},
		},
		// Serializes complex query maps
		"QueryComplexQueryMaps": {
			Params: &QueryMapsInput{
				ComplexMapArg: map[string]types.GreetingStruct{
					"bar": {
						Hi: ptr.String("Bar"),
					},
					"foo": {
						Hi: ptr.String("Foo"),
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&ComplexMapArg.entry.1.key=bar
			&ComplexMapArg.entry.1.value.hi=Bar
			&ComplexMapArg.entry.2.key=foo
			&ComplexMapArg.entry.2.value.hi=Foo`))
			},
		},
		// Does not serialize empty query maps
		"QueryEmptyQueryMaps": {
			Params: &QueryMapsInput{
				MapArg: map[string]string{},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08`))
			},
		},
		// Serializes query maps where the member has an xmlName trait
		"QueryQueryMapWithMemberXmlName": {
			Params: &QueryMapsInput{
				MapWithXmlMemberName: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&MapWithXmlMemberName.entry.1.K=bar
			&MapWithXmlMemberName.entry.1.V=Bar
			&MapWithXmlMemberName.entry.2.K=foo
			&MapWithXmlMemberName.entry.2.V=Foo`))
			},
		},
		// Serializes flattened query maps
		"QueryFlattenedQueryMaps": {
			Params: &QueryMapsInput{
				FlattenedMap: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&FlattenedMap.1.key=bar
			&FlattenedMap.1.value=Bar
			&FlattenedMap.2.key=foo
			&FlattenedMap.2.value=Foo`))
			},
		},
		// Serializes flattened query maps that use an xmlName
		"QueryFlattenedQueryMapsWithXmlName": {
			Params: &QueryMapsInput{
				FlattenedMapWithXmlName: map[string]string{
					"bar": "Bar",
					"foo": "Foo",
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&Hi.1.K=bar
			&Hi.1.V=Bar
			&Hi.2.K=foo
			&Hi.2.V=Foo`))
			},
		},
		// Serializes query map of lists
		"QueryQueryMapOfLists": {
			Params: &QueryMapsInput{
				MapOfLists: map[string][]string{
					"bar": {
						"C",
						"D",
					},
					"foo": {
						"A",
						"B",
					},
				},
			},
			ExpectMethod:  "POST",
			ExpectURIPath: "/",
			ExpectQuery:   []smithytesting.QueryItem{},
			ExpectHeader: http.Header{
				"Content-Type": []string{"application/x-www-form-urlencoded"},
			},
			BodyMediaType: "application/x-www-form-urlencoded",
			BodyAssert: func(actual io.Reader) error {
				return smithytesting.CompareURLFormReaderBytes(actual, []byte(`Action=QueryMaps
			&Version=2020-01-08
			&MapOfLists.entry.1.key=bar
			&MapOfLists.entry.1.value.member.1=C
			&MapOfLists.entry.1.value.member.2=D
			&MapOfLists.entry.2.key=foo
			&MapOfLists.entry.2.value.member.1=A
			&MapOfLists.entry.2.value.member.2=B`))
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			var actualReq *http.Request
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				actualReq = r.Clone(r.Context())
				if len(actualReq.URL.RawPath) == 0 {
					actualReq.URL.RawPath = actualReq.URL.Path
				}
				if v := actualReq.ContentLength; v != 0 {
					actualReq.Header.Set("Content-Length", strconv.FormatInt(v, 10))
				}
				var buf bytes.Buffer
				if _, err := io.Copy(&buf, r.Body); err != nil {
					t.Errorf("failed to read request body, %v", err)
				}
				actualReq.Body = ioutil.NopCloser(&buf)

				w.WriteHeader(200)
			}))
			defer server.Close()
			url := server.URL
			client := New(Options{
				APIOptions: []func(*middleware.Stack) error{
					func(s *middleware.Stack) error {
						s.Finalize.Clear()
						return nil
					},
				},
				EndpointResolver: EndpointResolverFunc(func(region string, options EndpointResolverOptions) (e aws.Endpoint, err error) {
					e.URL = url
					e.SigningRegion = "us-west-2"
					return e, err
				}),
				HTTPClient:               awshttp.NewBuildableClient(),
				IdempotencyTokenProvider: smithyrand.NewUUIDIdempotencyToken(&smithytesting.ByteLoop{}),
				Region:                   "us-west-2",
			})
			result, err := client.QueryMaps(context.Background(), c.Params)
			if err != nil {
				t.Fatalf("expect nil err, got %v", err)
			}
			if result == nil {
				t.Fatalf("expect not nil result")
			}
			if e, a := c.ExpectMethod, actualReq.Method; e != a {
				t.Errorf("expect %v method, got %v", e, a)
			}
			if e, a := c.ExpectURIPath, actualReq.URL.RawPath; e != a {
				t.Errorf("expect %v path, got %v", e, a)
			}
			queryItems := smithytesting.ParseRawQuery(actualReq.URL.RawQuery)
			smithytesting.AssertHasQuery(t, c.ExpectQuery, queryItems)
			smithytesting.AssertHasQueryKeys(t, c.RequireQuery, queryItems)
			smithytesting.AssertNotHaveQueryKeys(t, c.ForbidQuery, queryItems)
			smithytesting.AssertHasHeader(t, c.ExpectHeader, actualReq.Header)
			smithytesting.AssertHasHeaderKeys(t, c.RequireHeader, actualReq.Header)
			smithytesting.AssertNotHaveHeaderKeys(t, c.ForbidHeader, actualReq.Header)
			if c.BodyAssert != nil {
				if err := c.BodyAssert(actualReq.Body); err != nil {
					t.Errorf("expect body equal, got %v", err)
				}
			}
		})
	}
}
