// +build !ignore_autogenerated

// Code generated by deepcopy-gen. DO NOT EDIT.

package ignite

import (
	net "net"

	v1alpha1 "github.com/weaveworks/ignite/pkg/apis/meta/v1alpha1"
	pkgruntime "github.com/weaveworks/libgitops/pkg/runtime"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceVolume) DeepCopyInto(out *BlockDeviceVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceVolume.
func (in *BlockDeviceVolume) DeepCopy() *BlockDeviceVolume {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileMapping) DeepCopyInto(out *FileMapping) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileMapping.
func (in *FileMapping) DeepCopy() *FileMapping {
	if in == nil {
		return nil
	}
	out := new(FileMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Image) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	out.OCI = in.OCI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageStatus) DeepCopyInto(out *ImageStatus) {
	*out = *in
	in.OCISource.DeepCopyInto(&out.OCISource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageStatus.
func (in *ImageStatus) DeepCopy() *ImageStatus {
	if in == nil {
		return nil
	}
	out := new(ImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Kernel) DeepCopyInto(out *Kernel) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Kernel.
func (in *Kernel) DeepCopy() *Kernel {
	if in == nil {
		return nil
	}
	out := new(Kernel)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Kernel) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelSpec) DeepCopyInto(out *KernelSpec) {
	*out = *in
	out.OCI = in.OCI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelSpec.
func (in *KernelSpec) DeepCopy() *KernelSpec {
	if in == nil {
		return nil
	}
	out := new(KernelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KernelStatus) DeepCopyInto(out *KernelStatus) {
	*out = *in
	in.OCISource.DeepCopyInto(&out.OCISource)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KernelStatus.
func (in *KernelStatus) DeepCopy() *KernelStatus {
	if in == nil {
		return nil
	}
	out := new(KernelStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCIImageSource) DeepCopyInto(out *OCIImageSource) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(v1alpha1.OCIContentID)
		**out = **in
	}
	out.Size = in.Size
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCIImageSource.
func (in *OCIImageSource) DeepCopy() *OCIImageSource {
	if in == nil {
		return nil
	}
	out := new(OCIImageSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pool) DeepCopyInto(out *Pool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pool.
func (in *Pool) DeepCopy() *Pool {
	if in == nil {
		return nil
	}
	out := new(Pool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Pool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolDevice) DeepCopyInto(out *PoolDevice) {
	*out = *in
	out.Size = in.Size
	out.Parent = in.Parent
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolDevice.
func (in *PoolDevice) DeepCopy() *PoolDevice {
	if in == nil {
		return nil
	}
	out := new(PoolDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolSpec) DeepCopyInto(out *PoolSpec) {
	*out = *in
	out.MetadataSize = in.MetadataSize
	out.DataSize = in.DataSize
	out.AllocationSize = in.AllocationSize
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolSpec.
func (in *PoolSpec) DeepCopy() *PoolSpec {
	if in == nil {
		return nil
	}
	out := new(PoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PoolStatus) DeepCopyInto(out *PoolStatus) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*PoolDevice, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(PoolDevice)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PoolStatus.
func (in *PoolStatus) DeepCopy() *PoolStatus {
	if in == nil {
		return nil
	}
	out := new(PoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Runtime) DeepCopyInto(out *Runtime) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Runtime.
func (in *Runtime) DeepCopy() *Runtime {
	if in == nil {
		return nil
	}
	out := new(Runtime)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSH) DeepCopyInto(out *SSH) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSH.
func (in *SSH) DeepCopy() *SSH {
	if in == nil {
		return nil
	}
	out := new(SSH)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VM) DeepCopyInto(out *VM) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VM.
func (in *VM) DeepCopy() *VM {
	if in == nil {
		return nil
	}
	out := new(VM)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VM) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMImageSpec) DeepCopyInto(out *VMImageSpec) {
	*out = *in
	out.OCI = in.OCI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMImageSpec.
func (in *VMImageSpec) DeepCopy() *VMImageSpec {
	if in == nil {
		return nil
	}
	out := new(VMImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMKernelSpec) DeepCopyInto(out *VMKernelSpec) {
	*out = *in
	out.OCI = in.OCI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMKernelSpec.
func (in *VMKernelSpec) DeepCopy() *VMKernelSpec {
	if in == nil {
		return nil
	}
	out := new(VMKernelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMNetworkSpec) DeepCopyInto(out *VMNetworkSpec) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make(v1alpha1.PortMappings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMNetworkSpec.
func (in *VMNetworkSpec) DeepCopy() *VMNetworkSpec {
	if in == nil {
		return nil
	}
	out := new(VMNetworkSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSandboxSpec) DeepCopyInto(out *VMSandboxSpec) {
	*out = *in
	out.OCI = in.OCI
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSandboxSpec.
func (in *VMSandboxSpec) DeepCopy() *VMSandboxSpec {
	if in == nil {
		return nil
	}
	out := new(VMSandboxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMSpec) DeepCopyInto(out *VMSpec) {
	*out = *in
	out.Image = in.Image
	out.Sandbox = in.Sandbox
	out.Kernel = in.Kernel
	out.Memory = in.Memory
	out.DiskSize = in.DiskSize
	in.Network.DeepCopyInto(&out.Network)
	in.Storage.DeepCopyInto(&out.Storage)
	if in.CopyFiles != nil {
		in, out := &in.CopyFiles, &out.CopyFiles
		*out = make([]FileMapping, len(*in))
		copy(*out, *in)
	}
	if in.SSH != nil {
		in, out := &in.SSH, &out.SSH
		*out = new(SSH)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMSpec.
func (in *VMSpec) DeepCopy() *VMSpec {
	if in == nil {
		return nil
	}
	out := new(VMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStatus) DeepCopyInto(out *VMStatus) {
	*out = *in
	if in.Runtime != nil {
		in, out := &in.Runtime, &out.Runtime
		*out = new(Runtime)
		**out = **in
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = new(pkgruntime.Time)
		(*in).DeepCopyInto(*out)
	}
	if in.IPAddresses != nil {
		in, out := &in.IPAddresses, &out.IPAddresses
		*out = make(v1alpha1.IPAddresses, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = make(net.IP, len(*in))
				copy(*out, *in)
			}
		}
	}
	in.Image.DeepCopyInto(&out.Image)
	in.Kernel.DeepCopyInto(&out.Kernel)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStatus.
func (in *VMStatus) DeepCopy() *VMStatus {
	if in == nil {
		return nil
	}
	out := new(VMStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VMStorageSpec) DeepCopyInto(out *VMStorageSpec) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]VolumeMount, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VMStorageSpec.
func (in *VMStorageSpec) DeepCopy() *VMStorageSpec {
	if in == nil {
		return nil
	}
	out := new(VMStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Volume) DeepCopyInto(out *Volume) {
	*out = *in
	if in.BlockDevice != nil {
		in, out := &in.BlockDevice, &out.BlockDevice
		*out = new(BlockDeviceVolume)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Volume.
func (in *Volume) DeepCopy() *Volume {
	if in == nil {
		return nil
	}
	out := new(Volume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VolumeMount) DeepCopyInto(out *VolumeMount) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VolumeMount.
func (in *VolumeMount) DeepCopy() *VolumeMount {
	if in == nil {
		return nil
	}
	out := new(VolumeMount)
	in.DeepCopyInto(out)
	return out
}
