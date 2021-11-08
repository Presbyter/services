//go:generate truss -v --svcout ../src/service/auth-service auth.proto
//go:generate truss -v --svcout ../src/service/im-service im.proto
//go:generate truss -v --svcout ../src/service/gateway-service api.proto
package protobuf
