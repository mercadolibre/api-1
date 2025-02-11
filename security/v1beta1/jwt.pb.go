// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: security/v1beta1/jwt.proto

package v1beta1

import (
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	_ "istio.io/gogo-genproto/googleapis/google/api"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// JSON Web Token (JWT) token format for authentication as defined by
// [RFC 7519](https://tools.ietf.org/html/rfc7519). See [OAuth 2.0](https://tools.ietf.org/html/rfc6749) and
// [OIDC 1.0](http://openid.net/connect) for how this is used in the whole
// authentication flow.
//
// Examples:
//
// Spec for a JWT that is issued by `https://example.com`, with the audience claims must be either
// `bookstore_android.apps.example.com` or `bookstore_web.apps.example.com`.
// The token should be presented at the `Authorization` header (default). The JSON Web Key Set (JWKS)
// will be discovered following OpenID Connect protocol.
//
// ```yaml
// issuer: https://example.com
// audiences:
// - bookstore_android.apps.example.com
//   bookstore_web.apps.example.com
// ```
//
// This example specifies a token in a non-default location (`x-goog-iap-jwt-assertion` header). It also
// defines the URI to fetch JWKS explicitly.
//
// ```yaml
// issuer: https://example.com
// jwksUri: https://example.com/.secret/jwks.json
// fromHeaders:
// - "x-goog-iap-jwt-assertion"
// ```
type JWTRule struct {
	// Identifies the issuer that issued the JWT. See
	// [issuer](https://tools.ietf.org/html/rfc7519#section-4.1.1)
	// A JWT with different `iss` claim will be rejected.
	//
	// Example: https://foobar.auth0.com
	// Example: 1234567-compute@developer.gserviceaccount.com
	Issuer string `protobuf:"bytes,1,opt,name=issuer,proto3" json:"issuer,omitempty"`
	// The list of JWT
	// [audiences](https://tools.ietf.org/html/rfc7519#section-4.1.3).
	// that are allowed to access. A JWT containing any of these
	// audiences will be accepted.
	//
	// The service name will be accepted if audiences is empty.
	//
	// Example:
	//
	// ```yaml
	// audiences:
	// - bookstore_android.apps.example.com
	//   bookstore_web.apps.example.com
	// ```
	Audiences []string `protobuf:"bytes,2,rep,name=audiences,proto3" json:"audiences,omitempty"`
	// URL of the provider's public key set to validate signature of the
	// JWT. See [OpenID Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html#ProviderMetadata).
	//
	// Optional if the key set document can either (a) be retrieved from
	// [OpenID
	// Discovery](https://openid.net/specs/openid-connect-discovery-1_0.html) of
	// the issuer or (b) inferred from the email domain of the issuer (e.g. a
	// Google service account).
	//
	// Example: `https://www.googleapis.com/oauth2/v1/certs`
	//
	// Note: Only one of `jwksUri` and `jwks` should be used.
	JwksUri string `protobuf:"bytes,3,opt,name=jwks_uri,json=jwksUri,proto3" json:"jwks_uri,omitempty"`
	// JSON Web Key Set of public keys to validate signature of the JWT.
	// See https://auth0.com/docs/jwks.
	//
	// Note: Only one of `jwksUri` and `jwks` should be used.
	Jwks string `protobuf:"bytes,10,opt,name=jwks,proto3" json:"jwks,omitempty"`
	// List of header locations from which JWT is expected. For example, below is the location spec
	// if JWT is expected to be found in `x-jwt-assertion` header, and have "Bearer " prefix:
	//
	// ```yaml
	//   fromHeaders:
	//   - name: x-jwt-assertion
	//     prefix: "Bearer "
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromHeaders []*JWTHeader `protobuf:"bytes,6,rep,name=from_headers,json=fromHeaders,proto3" json:"from_headers,omitempty"`
	// List of query parameters from which JWT is expected. For example, if JWT is provided via query
	// parameter `my_token` (e.g /path?my_token=<JWT>), the config is:
	//
	// ```yaml
	//   fromParams:
	//   - "my_token"
	// ```
	//
	// Note: Requests with multiple tokens (at different locations) are not supported, the output principal of
	// such requests is undefined.
	FromParams []string `protobuf:"bytes,7,rep,name=from_params,json=fromParams,proto3" json:"from_params,omitempty"`
	// This field specifies the header name to output a successfully verified JWT payload to the
	// backend. The forwarded data is `base64_encoded(jwt_payload_in_JSON)`. If it is not specified,
	// the payload will not be emitted.
	OutputPayloadToHeader string `protobuf:"bytes,8,opt,name=output_payload_to_header,json=outputPayloadToHeader,proto3" json:"output_payload_to_header,omitempty"`
	// If set to true, the original token will be kept for the upstream request. Default is false.
	ForwardOriginalToken bool     `protobuf:"varint,9,opt,name=forward_original_token,json=forwardOriginalToken,proto3" json:"forward_original_token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTRule) Reset()         { *m = JWTRule{} }
func (m *JWTRule) String() string { return proto.CompactTextString(m) }
func (*JWTRule) ProtoMessage()    {}
func (*JWTRule) Descriptor() ([]byte, []int) {
	return fileDescriptor_163ab6fd32fb6b15, []int{0}
}
func (m *JWTRule) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTRule) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTRule.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWTRule) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTRule.Merge(m, src)
}
func (m *JWTRule) XXX_Size() int {
	return m.Size()
}
func (m *JWTRule) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTRule.DiscardUnknown(m)
}

var xxx_messageInfo_JWTRule proto.InternalMessageInfo

func (m *JWTRule) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *JWTRule) GetAudiences() []string {
	if m != nil {
		return m.Audiences
	}
	return nil
}

func (m *JWTRule) GetJwksUri() string {
	if m != nil {
		return m.JwksUri
	}
	return ""
}

func (m *JWTRule) GetJwks() string {
	if m != nil {
		return m.Jwks
	}
	return ""
}

func (m *JWTRule) GetFromHeaders() []*JWTHeader {
	if m != nil {
		return m.FromHeaders
	}
	return nil
}

func (m *JWTRule) GetFromParams() []string {
	if m != nil {
		return m.FromParams
	}
	return nil
}

func (m *JWTRule) GetOutputPayloadToHeader() string {
	if m != nil {
		return m.OutputPayloadToHeader
	}
	return ""
}

func (m *JWTRule) GetForwardOriginalToken() bool {
	if m != nil {
		return m.ForwardOriginalToken
	}
	return false
}

// This message specifies a header location to extract JWT token.
type JWTHeader struct {
	// The HTTP header name.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// The prefix that should be stripped before decoding the token.
	// For example, for "Authorization: Bearer <token>", prefix="Bearer " with a space at the end.
	// If the header doesn't have this exact prefix, it is considered invalid.
	Prefix               string   `protobuf:"bytes,2,opt,name=prefix,proto3" json:"prefix,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTHeader) Reset()         { *m = JWTHeader{} }
func (m *JWTHeader) String() string { return proto.CompactTextString(m) }
func (*JWTHeader) ProtoMessage()    {}
func (*JWTHeader) Descriptor() ([]byte, []int) {
	return fileDescriptor_163ab6fd32fb6b15, []int{1}
}
func (m *JWTHeader) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTHeader) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTHeader.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *JWTHeader) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTHeader.Merge(m, src)
}
func (m *JWTHeader) XXX_Size() int {
	return m.Size()
}
func (m *JWTHeader) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTHeader.DiscardUnknown(m)
}

var xxx_messageInfo_JWTHeader proto.InternalMessageInfo

func (m *JWTHeader) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *JWTHeader) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

func init() {
	proto.RegisterType((*JWTRule)(nil), "istio.security.v1beta1.JWTRule")
	proto.RegisterType((*JWTHeader)(nil), "istio.security.v1beta1.JWTHeader")
}

func init() { proto.RegisterFile("security/v1beta1/jwt.proto", fileDescriptor_163ab6fd32fb6b15) }

var fileDescriptor_163ab6fd32fb6b15 = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xc1, 0x4e, 0xea, 0x40,
	0x14, 0x86, 0xd3, 0x42, 0x0a, 0x1d, 0xee, 0x6a, 0x72, 0x2f, 0x99, 0x4b, 0x10, 0x2a, 0xab, 0x26,
	0x26, 0x6d, 0x50, 0x13, 0x57, 0x2e, 0x24, 0x2e, 0x0c, 0x1b, 0x49, 0x53, 0x43, 0xe2, 0x66, 0x32,
	0xd0, 0x29, 0x0c, 0x94, 0x4e, 0x33, 0x33, 0x05, 0x79, 0x43, 0xe3, 0xca, 0x47, 0x30, 0x7d, 0x12,
	0xd3, 0x69, 0xd5, 0x68, 0xdc, 0x9d, 0xf3, 0xfd, 0xe7, 0xb4, 0xff, 0x9c, 0x1f, 0xf4, 0x24, 0x5d,
	0xe6, 0x82, 0xa9, 0xa3, 0xbf, 0x1f, 0x2f, 0xa8, 0x22, 0x63, 0x7f, 0x73, 0x50, 0x5e, 0x26, 0xb8,
	0xe2, 0xb0, 0xcb, 0xa4, 0x62, 0xdc, 0xfb, 0x98, 0xf0, 0xea, 0x89, 0xde, 0x70, 0xc5, 0xf9, 0x2a,
	0xa1, 0x3e, 0xc9, 0x98, 0x1f, 0x33, 0x9a, 0x44, 0x78, 0x41, 0xd7, 0x64, 0xcf, 0xb8, 0xa8, 0x16,
	0x47, 0x2f, 0x26, 0x68, 0x4d, 0xe7, 0x61, 0x90, 0x27, 0x14, 0xf6, 0x81, 0xc5, 0xa4, 0xcc, 0xa9,
	0x40, 0x86, 0x63, 0xb8, 0xf6, 0xa4, 0x59, 0xdc, 0x18, 0x66, 0x50, 0x33, 0xd8, 0x07, 0x36, 0xc9,
	0x23, 0x46, 0xd3, 0x25, 0x95, 0xc8, 0x74, 0x1a, 0xae, 0x1d, 0x7c, 0x01, 0xf8, 0x1f, 0xb4, 0x37,
	0x87, 0xad, 0xc4, 0xb9, 0x60, 0xa8, 0x51, 0x6e, 0x07, 0xad, 0xb2, 0x7f, 0x10, 0x0c, 0x42, 0xd0,
	0x2c, 0x4b, 0x04, 0x34, 0xd6, 0x35, 0xbc, 0x05, 0x7f, 0x62, 0xc1, 0x77, 0x78, 0x4d, 0x49, 0x44,
	0x85, 0x44, 0x96, 0xd3, 0x70, 0x3b, 0xe7, 0xa7, 0xde, 0xef, 0xcf, 0xf0, 0xa6, 0xf3, 0xf0, 0x4e,
	0x4f, 0x06, 0x9d, 0x72, 0xad, 0xaa, 0x25, 0x1c, 0x02, 0xdd, 0xe2, 0x8c, 0x08, 0xb2, 0x93, 0xa8,
	0xa5, 0x4d, 0x81, 0x12, 0xcd, 0x34, 0x81, 0x57, 0x00, 0xf1, 0x5c, 0x65, 0xb9, 0xc2, 0x19, 0x39,
	0x26, 0x9c, 0x44, 0x58, 0xf1, 0xfa, 0x9f, 0xa8, 0xad, 0xed, 0xfc, 0xab, 0xf4, 0x59, 0x25, 0x87,
	0xbc, 0xfa, 0x34, 0xbc, 0x04, 0xdd, 0x98, 0x8b, 0x03, 0x11, 0x11, 0xe6, 0x82, 0xad, 0x58, 0x4a,
	0x12, 0xac, 0xf8, 0x96, 0xa6, 0xc8, 0x76, 0x0c, 0xb7, 0x1d, 0xfc, 0xad, 0xd5, 0xfb, 0x5a, 0x0c,
	0x4b, 0x6d, 0x74, 0x0d, 0xec, 0x4f, 0xa7, 0x10, 0x81, 0x66, 0x4a, 0x76, 0xf4, 0xdb, 0x2d, 0x35,
	0x81, 0x5d, 0x60, 0x65, 0x82, 0xc6, 0xec, 0x09, 0x99, 0xda, 0x43, 0xdd, 0x4d, 0xce, 0x9e, 0x8b,
	0x81, 0xf1, 0x5a, 0x0c, 0x8c, 0xb7, 0x62, 0x60, 0x3c, 0x9e, 0x54, 0xb7, 0x60, 0x5c, 0x87, 0xf7,
	0x33, 0xfb, 0x85, 0xa5, 0xf3, 0xbb, 0x78, 0x0f, 0x00, 0x00, 0xff, 0xff, 0xd0, 0x38, 0x7c, 0xfc,
	0x16, 0x02, 0x00, 0x00,
}

func (m *JWTRule) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTRule) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWTRule) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Jwks) > 0 {
		i -= len(m.Jwks)
		copy(dAtA[i:], m.Jwks)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Jwks)))
		i--
		dAtA[i] = 0x52
	}
	if m.ForwardOriginalToken {
		i--
		if m.ForwardOriginalToken {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x48
	}
	if len(m.OutputPayloadToHeader) > 0 {
		i -= len(m.OutputPayloadToHeader)
		copy(dAtA[i:], m.OutputPayloadToHeader)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.OutputPayloadToHeader)))
		i--
		dAtA[i] = 0x42
	}
	if len(m.FromParams) > 0 {
		for iNdEx := len(m.FromParams) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.FromParams[iNdEx])
			copy(dAtA[i:], m.FromParams[iNdEx])
			i = encodeVarintJwt(dAtA, i, uint64(len(m.FromParams[iNdEx])))
			i--
			dAtA[i] = 0x3a
		}
	}
	if len(m.FromHeaders) > 0 {
		for iNdEx := len(m.FromHeaders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.FromHeaders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintJwt(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.JwksUri) > 0 {
		i -= len(m.JwksUri)
		copy(dAtA[i:], m.JwksUri)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.JwksUri)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Audiences) > 0 {
		for iNdEx := len(m.Audiences) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Audiences[iNdEx])
			copy(dAtA[i:], m.Audiences[iNdEx])
			i = encodeVarintJwt(dAtA, i, uint64(len(m.Audiences[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *JWTHeader) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTHeader) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *JWTHeader) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Prefix) > 0 {
		i -= len(m.Prefix)
		copy(dAtA[i:], m.Prefix)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Prefix)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintJwt(dAtA []byte, offset int, v uint64) int {
	offset -= sovJwt(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *JWTRule) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if len(m.Audiences) > 0 {
		for _, s := range m.Audiences {
			l = len(s)
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	l = len(m.JwksUri)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if len(m.FromHeaders) > 0 {
		for _, e := range m.FromHeaders {
			l = e.Size()
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	if len(m.FromParams) > 0 {
		for _, s := range m.FromParams {
			l = len(s)
			n += 1 + l + sovJwt(uint64(l))
		}
	}
	l = len(m.OutputPayloadToHeader)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.ForwardOriginalToken {
		n += 2
	}
	l = len(m.Jwks)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *JWTHeader) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	l = len(m.Prefix)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovJwt(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozJwt(x uint64) (n int) {
	return sovJwt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *JWTRule) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JWTRule: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTRule: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Audiences", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Audiences = append(m.Audiences, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field JwksUri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.JwksUri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromHeaders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromHeaders = append(m.FromHeaders, &JWTHeader{})
			if err := m.FromHeaders[len(m.FromHeaders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field FromParams", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.FromParams = append(m.FromParams, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field OutputPayloadToHeader", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.OutputPayloadToHeader = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ForwardOriginalToken", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ForwardOriginalToken = bool(v != 0)
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Jwks", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Jwks = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *JWTHeader) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JWTHeader: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTHeader: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prefix", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthJwt
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Prefix = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthJwt
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipJwt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthJwt
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupJwt
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthJwt
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthJwt        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJwt          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupJwt = fmt.Errorf("proto: unexpected end of group")
)
