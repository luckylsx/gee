package gee

import (
	"reflect"
	"testing"
)

func Test_newRouter(t *testing.T) {
	tests := []struct {
		name string
		want *router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := newRouter(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newRouter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parsePattern(t *testing.T) {
	parts := parsePattern("/hello/:name")
	t.Logf("parts =%v", parts)

	parts1 := parsePattern("/")
	t.Logf("parts1 =%v", parts1)
	// type args struct {
	// 	pattern string
	// }
	// tests := []struct {
	// 	name string
	// 	args args
	// 	want []string
	// }{
	// 	// TODO: Add test cases.
	// }
	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {
	// 		if got := parsePattern(tt.args.pattern); !reflect.DeepEqual(got, tt.want) {
	// 			t.Errorf("parsePattern() = %v, want %v", got, tt.want)
	// 		}
	// 	})
	// }
}

func Test_router_addRouter(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method  string
		pattern string
		handler HandlerFunc
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			r.addRouter(tt.args.method, tt.args.pattern, tt.args.handler)
		})
	}
}

func Test_router_getRouter(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		method string
		path   string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *node
		want1  map[string]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			got, got1 := r.getRouter(tt.args.method, tt.args.path)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("router.getRouter() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("router.getRouter() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_router_handle(t *testing.T) {
	type fields struct {
		roots    map[string]*node
		handlers map[string]HandlerFunc
	}
	type args struct {
		c *Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &router{
				roots:    tt.fields.roots,
				handlers: tt.fields.handlers,
			}
			r.handle(tt.args.c)
		})
	}
}
