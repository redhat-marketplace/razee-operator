// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APIKeyRef) DeepCopyInto(out *APIKeyRef) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APIKeyRef.
func (in *APIKeyRef) DeepCopy() *APIKeyRef {
	if in == nil {
		return nil
	}
	out := new(APIKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccesKeyIDRef) DeepCopyInto(out *AccesKeyIDRef) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccesKeyIDRef.
func (in *AccesKeyIDRef) DeepCopy() *AccesKeyIDRef {
	if in == nil {
		return nil
	}
	out := new(AccesKeyIDRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Auth) DeepCopyInto(out *Auth) {
	*out = *in
	if in.Hmac != nil {
		in, out := &in.Hmac, &out.Hmac
		*out = new(Hmac)
		(*in).DeepCopyInto(*out)
	}
	if in.Iam != nil {
		in, out := &in.Iam, &out.Iam
		*out = new(Iam)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Auth.
func (in *Auth) DeepCopy() *Auth {
	if in == nil {
		return nil
	}
	out := new(Auth)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Header) DeepCopyInto(out *Header) {
	{
		in := &in
		*out = make(Header, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Header.
func (in Header) DeepCopy() Header {
	if in == nil {
		return nil
	}
	out := new(Header)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Hmac) DeepCopyInto(out *Hmac) {
	*out = *in
	in.AccesKeyIDRef.DeepCopyInto(&out.AccesKeyIDRef)
	in.SecretAccessKeyRef.DeepCopyInto(&out.SecretAccessKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Hmac.
func (in *Hmac) DeepCopy() *Hmac {
	if in == nil {
		return nil
	}
	out := new(Hmac)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Iam) DeepCopyInto(out *Iam) {
	*out = *in
	in.APIKeyRef.DeepCopyInto(&out.APIKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Iam.
func (in *Iam) DeepCopy() *Iam {
	if in == nil {
		return nil
	}
	out := new(Iam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Items) DeepCopyInto(out *Items) {
	*out = *in
	in.Options.DeepCopyInto(&out.Options)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Items.
func (in *Items) DeepCopy() *Items {
	if in == nil {
		return nil
	}
	out := new(Items)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Log) DeepCopyInto(out *Log) {
	{
		in := &in
		*out = make(Log, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
		return
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Log.
func (in Log) DeepCopy() Log {
	if in == nil {
		return nil
	}
	out := new(Log)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Options) DeepCopyInto(out *Options) {
	*out = *in
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make(map[string]Header, len(*in))
		for key, val := range *in {
			var outVal map[string]string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(Header, len(*in))
				for key, val := range *in {
					(*out)[key] = val
				}
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Options.
func (in *Options) DeepCopy() *Options {
	if in == nil {
		return nil
	}
	out := new(Options)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RazeeLogs) DeepCopyInto(out *RazeeLogs) {
	*out = *in
	if in.Log != nil {
		in, out := &in.Log, &out.Log
		*out = make(Log, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RazeeLogs.
func (in *RazeeLogs) DeepCopy() *RazeeLogs {
	if in == nil {
		return nil
	}
	out := new(RazeeLogs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteResourceS3) DeepCopyInto(out *RemoteResourceS3) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteResourceS3.
func (in *RemoteResourceS3) DeepCopy() *RemoteResourceS3 {
	if in == nil {
		return nil
	}
	out := new(RemoteResourceS3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteResourceS3) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteResourceS3List) DeepCopyInto(out *RemoteResourceS3List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RemoteResourceS3, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteResourceS3List.
func (in *RemoteResourceS3List) DeepCopy() *RemoteResourceS3List {
	if in == nil {
		return nil
	}
	out := new(RemoteResourceS3List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RemoteResourceS3List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteResourceS3Spec) DeepCopyInto(out *RemoteResourceS3Spec) {
	*out = *in
	in.Auth.DeepCopyInto(&out.Auth)
	if in.Requests != nil {
		in, out := &in.Requests, &out.Requests
		*out = make([]Items, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteResourceS3Spec.
func (in *RemoteResourceS3Spec) DeepCopy() *RemoteResourceS3Spec {
	if in == nil {
		return nil
	}
	out := new(RemoteResourceS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteResourceS3Status) DeepCopyInto(out *RemoteResourceS3Status) {
	*out = *in
	in.RazeeLogs.DeepCopyInto(&out.RazeeLogs)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteResourceS3Status.
func (in *RemoteResourceS3Status) DeepCopy() *RemoteResourceS3Status {
	if in == nil {
		return nil
	}
	out := new(RemoteResourceS3Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretAccessKeyRef) DeepCopyInto(out *SecretAccessKeyRef) {
	*out = *in
	in.ValueFrom.DeepCopyInto(&out.ValueFrom)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretAccessKeyRef.
func (in *SecretAccessKeyRef) DeepCopy() *SecretAccessKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretAccessKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFrom) DeepCopyInto(out *ValueFrom) {
	*out = *in
	in.SecretKeyRef.DeepCopyInto(&out.SecretKeyRef)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFrom.
func (in *ValueFrom) DeepCopy() *ValueFrom {
	if in == nil {
		return nil
	}
	out := new(ValueFrom)
	in.DeepCopyInto(out)
	return out
}
