// +build !ignore_autogenerated

/*
MIT License

Copyright (c) 2020 RH France Solution Architects

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FruitsCatalogG1) DeepCopyInto(out *FruitsCatalogG1) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FruitsCatalogG1.
func (in *FruitsCatalogG1) DeepCopy() *FruitsCatalogG1 {
	if in == nil {
		return nil
	}
	out := new(FruitsCatalogG1)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FruitsCatalogG1) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FruitsCatalogG1List) DeepCopyInto(out *FruitsCatalogG1List) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]FruitsCatalogG1, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FruitsCatalogG1List.
func (in *FruitsCatalogG1List) DeepCopy() *FruitsCatalogG1List {
	if in == nil {
		return nil
	}
	out := new(FruitsCatalogG1List)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FruitsCatalogG1List) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FruitsCatalogG1Spec) DeepCopyInto(out *FruitsCatalogG1Spec) {
	*out = *in
	out.WebApp = in.WebApp
	out.MongoDB = in.MongoDB
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FruitsCatalogG1Spec.
func (in *FruitsCatalogG1Spec) DeepCopy() *FruitsCatalogG1Spec {
	if in == nil {
		return nil
	}
	out := new(FruitsCatalogG1Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FruitsCatalogG1Status) DeepCopyInto(out *FruitsCatalogG1Status) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FruitsCatalogG1Status.
func (in *FruitsCatalogG1Status) DeepCopy() *FruitsCatalogG1Status {
	if in == nil {
		return nil
	}
	out := new(FruitsCatalogG1Status)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSpec) DeepCopyInto(out *IngressSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSpec.
func (in *IngressSpec) DeepCopy() *IngressSpec {
	if in == nil {
		return nil
	}
	out := new(IngressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MongoDBSpec) DeepCopyInto(out *MongoDBSpec) {
	*out = *in
	out.SecretRef = in.SecretRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MongoDBSpec.
func (in *MongoDBSpec) DeepCopy() *MongoDBSpec {
	if in == nil {
		return nil
	}
	out := new(MongoDBSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretRefSpec) DeepCopyInto(out *SecretRefSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretRefSpec.
func (in *SecretRefSpec) DeepCopy() *SecretRefSpec {
	if in == nil {
		return nil
	}
	out := new(SecretRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebAppSpec) DeepCopyInto(out *WebAppSpec) {
	*out = *in
	out.Ingress = in.Ingress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebAppSpec.
func (in *WebAppSpec) DeepCopy() *WebAppSpec {
	if in == nil {
		return nil
	}
	out := new(WebAppSpec)
	in.DeepCopyInto(out)
	return out
}
