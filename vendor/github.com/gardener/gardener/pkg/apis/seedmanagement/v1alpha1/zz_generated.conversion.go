//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright (c) SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	core "github.com/gardener/gardener/pkg/apis/core"
	v1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	seedmanagement "github.com/gardener/gardener/pkg/apis/seedmanagement"
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*GardenletDeployment)(nil), (*seedmanagement.GardenletDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_GardenletDeployment_To_seedmanagement_GardenletDeployment(a.(*GardenletDeployment), b.(*seedmanagement.GardenletDeployment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.GardenletDeployment)(nil), (*GardenletDeployment)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_GardenletDeployment_To_v1alpha1_GardenletDeployment(a.(*seedmanagement.GardenletDeployment), b.(*GardenletDeployment), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Image)(nil), (*seedmanagement.Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Image_To_seedmanagement_Image(a.(*Image), b.(*seedmanagement.Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.Image)(nil), (*Image)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_Image_To_v1alpha1_Image(a.(*seedmanagement.Image), b.(*Image), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeed)(nil), (*seedmanagement.ManagedSeed)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed(a.(*ManagedSeed), b.(*seedmanagement.ManagedSeed), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeed)(nil), (*ManagedSeed)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed(a.(*seedmanagement.ManagedSeed), b.(*ManagedSeed), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedList)(nil), (*seedmanagement.ManagedSeedList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedList_To_seedmanagement_ManagedSeedList(a.(*ManagedSeedList), b.(*seedmanagement.ManagedSeedList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedList)(nil), (*ManagedSeedList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedList_To_v1alpha1_ManagedSeedList(a.(*seedmanagement.ManagedSeedList), b.(*ManagedSeedList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedSet)(nil), (*seedmanagement.ManagedSeedSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet(a.(*ManagedSeedSet), b.(*seedmanagement.ManagedSeedSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedSet)(nil), (*ManagedSeedSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet(a.(*seedmanagement.ManagedSeedSet), b.(*ManagedSeedSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedSetList)(nil), (*seedmanagement.ManagedSeedSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedSetList_To_seedmanagement_ManagedSeedSetList(a.(*ManagedSeedSetList), b.(*seedmanagement.ManagedSeedSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedSetList)(nil), (*ManagedSeedSetList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedSetList_To_v1alpha1_ManagedSeedSetList(a.(*seedmanagement.ManagedSeedSetList), b.(*ManagedSeedSetList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedSetSpec)(nil), (*seedmanagement.ManagedSeedSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec(a.(*ManagedSeedSetSpec), b.(*seedmanagement.ManagedSeedSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedSetSpec)(nil), (*ManagedSeedSetSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec(a.(*seedmanagement.ManagedSeedSetSpec), b.(*ManagedSeedSetSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedSetStatus)(nil), (*seedmanagement.ManagedSeedSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus(a.(*ManagedSeedSetStatus), b.(*seedmanagement.ManagedSeedSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedSetStatus)(nil), (*ManagedSeedSetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus(a.(*seedmanagement.ManagedSeedSetStatus), b.(*ManagedSeedSetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedSpec)(nil), (*seedmanagement.ManagedSeedSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(a.(*ManagedSeedSpec), b.(*seedmanagement.ManagedSeedSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedSpec)(nil), (*ManagedSeedSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(a.(*seedmanagement.ManagedSeedSpec), b.(*ManagedSeedSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedStatus)(nil), (*seedmanagement.ManagedSeedStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus(a.(*ManagedSeedStatus), b.(*seedmanagement.ManagedSeedStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedStatus)(nil), (*ManagedSeedStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus(a.(*seedmanagement.ManagedSeedStatus), b.(*ManagedSeedStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ManagedSeedTemplate)(nil), (*seedmanagement.ManagedSeedTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate(a.(*ManagedSeedTemplate), b.(*seedmanagement.ManagedSeedTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.ManagedSeedTemplate)(nil), (*ManagedSeedTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate(a.(*seedmanagement.ManagedSeedTemplate), b.(*ManagedSeedTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PendingReplica)(nil), (*seedmanagement.PendingReplica)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_PendingReplica_To_seedmanagement_PendingReplica(a.(*PendingReplica), b.(*seedmanagement.PendingReplica), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.PendingReplica)(nil), (*PendingReplica)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_PendingReplica_To_v1alpha1_PendingReplica(a.(*seedmanagement.PendingReplica), b.(*PendingReplica), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RollingUpdateStrategy)(nil), (*seedmanagement.RollingUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RollingUpdateStrategy_To_seedmanagement_RollingUpdateStrategy(a.(*RollingUpdateStrategy), b.(*seedmanagement.RollingUpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.RollingUpdateStrategy)(nil), (*RollingUpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_RollingUpdateStrategy_To_v1alpha1_RollingUpdateStrategy(a.(*seedmanagement.RollingUpdateStrategy), b.(*RollingUpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Shoot)(nil), (*seedmanagement.Shoot)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Shoot_To_seedmanagement_Shoot(a.(*Shoot), b.(*seedmanagement.Shoot), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.Shoot)(nil), (*Shoot)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_Shoot_To_v1alpha1_Shoot(a.(*seedmanagement.Shoot), b.(*Shoot), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*UpdateStrategy)(nil), (*seedmanagement.UpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_UpdateStrategy_To_seedmanagement_UpdateStrategy(a.(*UpdateStrategy), b.(*seedmanagement.UpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*seedmanagement.UpdateStrategy)(nil), (*UpdateStrategy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_UpdateStrategy_To_v1alpha1_UpdateStrategy(a.(*seedmanagement.UpdateStrategy), b.(*UpdateStrategy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*core.SeedTemplate)(nil), (*v1beta1.SeedTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_SeedTemplate_To_v1beta1_SeedTemplate(a.(*core.SeedTemplate), b.(*v1beta1.SeedTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*core.ShootTemplate)(nil), (*v1beta1.ShootTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_core_ShootTemplate_To_v1beta1_ShootTemplate(a.(*core.ShootTemplate), b.(*v1beta1.ShootTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*seedmanagement.Gardenlet)(nil), (*Gardenlet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_seedmanagement_Gardenlet_To_v1alpha1_Gardenlet(a.(*seedmanagement.Gardenlet), b.(*Gardenlet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*Gardenlet)(nil), (*seedmanagement.Gardenlet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Gardenlet_To_seedmanagement_Gardenlet(a.(*Gardenlet), b.(*seedmanagement.Gardenlet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.SeedTemplate)(nil), (*core.SeedTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_SeedTemplate_To_core_SeedTemplate(a.(*v1beta1.SeedTemplate), b.(*core.SeedTemplate), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*v1beta1.ShootTemplate)(nil), (*core.ShootTemplate)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1beta1_ShootTemplate_To_core_ShootTemplate(a.(*v1beta1.ShootTemplate), b.(*core.ShootTemplate), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_Gardenlet_To_seedmanagement_Gardenlet(in *Gardenlet, out *seedmanagement.Gardenlet, s conversion.Scope) error {
	out.Deployment = (*seedmanagement.GardenletDeployment)(unsafe.Pointer(in.Deployment))
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Config, &out.Config, s); err != nil {
		return err
	}
	out.Bootstrap = (*seedmanagement.Bootstrap)(unsafe.Pointer(in.Bootstrap))
	out.MergeWithParent = (*bool)(unsafe.Pointer(in.MergeWithParent))
	return nil
}

func autoConvert_seedmanagement_Gardenlet_To_v1alpha1_Gardenlet(in *seedmanagement.Gardenlet, out *Gardenlet, s conversion.Scope) error {
	out.Deployment = (*GardenletDeployment)(unsafe.Pointer(in.Deployment))
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Config, &out.Config, s); err != nil {
		return err
	}
	out.Bootstrap = (*Bootstrap)(unsafe.Pointer(in.Bootstrap))
	out.MergeWithParent = (*bool)(unsafe.Pointer(in.MergeWithParent))
	return nil
}

func autoConvert_v1alpha1_GardenletDeployment_To_seedmanagement_GardenletDeployment(in *GardenletDeployment, out *seedmanagement.GardenletDeployment, s conversion.Scope) error {
	out.ReplicaCount = (*int32)(unsafe.Pointer(in.ReplicaCount))
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.ServiceAccountName = (*string)(unsafe.Pointer(in.ServiceAccountName))
	out.Image = (*seedmanagement.Image)(unsafe.Pointer(in.Image))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.PodLabels = *(*map[string]string)(unsafe.Pointer(&in.PodLabels))
	out.PodAnnotations = *(*map[string]string)(unsafe.Pointer(&in.PodAnnotations))
	out.AdditionalVolumes = *(*[]v1.Volume)(unsafe.Pointer(&in.AdditionalVolumes))
	out.AdditionalVolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.AdditionalVolumeMounts))
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.VPA = (*bool)(unsafe.Pointer(in.VPA))
	return nil
}

// Convert_v1alpha1_GardenletDeployment_To_seedmanagement_GardenletDeployment is an autogenerated conversion function.
func Convert_v1alpha1_GardenletDeployment_To_seedmanagement_GardenletDeployment(in *GardenletDeployment, out *seedmanagement.GardenletDeployment, s conversion.Scope) error {
	return autoConvert_v1alpha1_GardenletDeployment_To_seedmanagement_GardenletDeployment(in, out, s)
}

func autoConvert_seedmanagement_GardenletDeployment_To_v1alpha1_GardenletDeployment(in *seedmanagement.GardenletDeployment, out *GardenletDeployment, s conversion.Scope) error {
	out.ReplicaCount = (*int32)(unsafe.Pointer(in.ReplicaCount))
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	out.ServiceAccountName = (*string)(unsafe.Pointer(in.ServiceAccountName))
	out.Image = (*Image)(unsafe.Pointer(in.Image))
	out.Resources = (*v1.ResourceRequirements)(unsafe.Pointer(in.Resources))
	out.PodLabels = *(*map[string]string)(unsafe.Pointer(&in.PodLabels))
	out.PodAnnotations = *(*map[string]string)(unsafe.Pointer(&in.PodAnnotations))
	out.AdditionalVolumes = *(*[]v1.Volume)(unsafe.Pointer(&in.AdditionalVolumes))
	out.AdditionalVolumeMounts = *(*[]v1.VolumeMount)(unsafe.Pointer(&in.AdditionalVolumeMounts))
	out.Env = *(*[]v1.EnvVar)(unsafe.Pointer(&in.Env))
	out.VPA = (*bool)(unsafe.Pointer(in.VPA))
	return nil
}

// Convert_seedmanagement_GardenletDeployment_To_v1alpha1_GardenletDeployment is an autogenerated conversion function.
func Convert_seedmanagement_GardenletDeployment_To_v1alpha1_GardenletDeployment(in *seedmanagement.GardenletDeployment, out *GardenletDeployment, s conversion.Scope) error {
	return autoConvert_seedmanagement_GardenletDeployment_To_v1alpha1_GardenletDeployment(in, out, s)
}

func autoConvert_v1alpha1_Image_To_seedmanagement_Image(in *Image, out *seedmanagement.Image, s conversion.Scope) error {
	out.Repository = (*string)(unsafe.Pointer(in.Repository))
	out.Tag = (*string)(unsafe.Pointer(in.Tag))
	out.PullPolicy = (*v1.PullPolicy)(unsafe.Pointer(in.PullPolicy))
	return nil
}

// Convert_v1alpha1_Image_To_seedmanagement_Image is an autogenerated conversion function.
func Convert_v1alpha1_Image_To_seedmanagement_Image(in *Image, out *seedmanagement.Image, s conversion.Scope) error {
	return autoConvert_v1alpha1_Image_To_seedmanagement_Image(in, out, s)
}

func autoConvert_seedmanagement_Image_To_v1alpha1_Image(in *seedmanagement.Image, out *Image, s conversion.Scope) error {
	out.Repository = (*string)(unsafe.Pointer(in.Repository))
	out.Tag = (*string)(unsafe.Pointer(in.Tag))
	out.PullPolicy = (*v1.PullPolicy)(unsafe.Pointer(in.PullPolicy))
	return nil
}

// Convert_seedmanagement_Image_To_v1alpha1_Image is an autogenerated conversion function.
func Convert_seedmanagement_Image_To_v1alpha1_Image(in *seedmanagement.Image, out *Image, s conversion.Scope) error {
	return autoConvert_seedmanagement_Image_To_v1alpha1_Image(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed(in *ManagedSeed, out *seedmanagement.ManagedSeed, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed(in *ManagedSeed, out *seedmanagement.ManagedSeed, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed(in *seedmanagement.ManagedSeed, out *ManagedSeed, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed(in *seedmanagement.ManagedSeed, out *ManagedSeed, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedList_To_seedmanagement_ManagedSeedList(in *ManagedSeedList, out *seedmanagement.ManagedSeedList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]seedmanagement.ManagedSeed, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ManagedSeed_To_seedmanagement_ManagedSeed(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ManagedSeedList_To_seedmanagement_ManagedSeedList is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedList_To_seedmanagement_ManagedSeedList(in *ManagedSeedList, out *seedmanagement.ManagedSeedList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedList_To_seedmanagement_ManagedSeedList(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedList_To_v1alpha1_ManagedSeedList(in *seedmanagement.ManagedSeedList, out *ManagedSeedList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedSeed, len(*in))
		for i := range *in {
			if err := Convert_seedmanagement_ManagedSeed_To_v1alpha1_ManagedSeed(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_seedmanagement_ManagedSeedList_To_v1alpha1_ManagedSeedList is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedList_To_v1alpha1_ManagedSeedList(in *seedmanagement.ManagedSeedList, out *ManagedSeedList, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedList_To_v1alpha1_ManagedSeedList(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet(in *ManagedSeedSet, out *seedmanagement.ManagedSeedSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet(in *ManagedSeedSet, out *seedmanagement.ManagedSeedSet, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet(in *seedmanagement.ManagedSeedSet, out *ManagedSeedSet, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet(in *seedmanagement.ManagedSeedSet, out *ManagedSeedSet, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedSetList_To_seedmanagement_ManagedSeedSetList(in *ManagedSeedSetList, out *seedmanagement.ManagedSeedSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]seedmanagement.ManagedSeedSet, len(*in))
		for i := range *in {
			if err := Convert_v1alpha1_ManagedSeedSet_To_seedmanagement_ManagedSeedSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_v1alpha1_ManagedSeedSetList_To_seedmanagement_ManagedSeedSetList is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedSetList_To_seedmanagement_ManagedSeedSetList(in *ManagedSeedSetList, out *seedmanagement.ManagedSeedSetList, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedSetList_To_seedmanagement_ManagedSeedSetList(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedSetList_To_v1alpha1_ManagedSeedSetList(in *seedmanagement.ManagedSeedSetList, out *ManagedSeedSetList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ManagedSeedSet, len(*in))
		for i := range *in {
			if err := Convert_seedmanagement_ManagedSeedSet_To_v1alpha1_ManagedSeedSet(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

// Convert_seedmanagement_ManagedSeedSetList_To_v1alpha1_ManagedSeedSetList is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedSetList_To_v1alpha1_ManagedSeedSetList(in *seedmanagement.ManagedSeedSetList, out *ManagedSeedSetList, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedSetList_To_v1alpha1_ManagedSeedSetList(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec(in *ManagedSeedSetSpec, out *seedmanagement.ManagedSeedSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = in.Selector
	if err := Convert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_v1beta1_ShootTemplate_To_core_ShootTemplate(&in.ShootTemplate, &out.ShootTemplate, s); err != nil {
		return err
	}
	out.UpdateStrategy = (*seedmanagement.UpdateStrategy)(unsafe.Pointer(in.UpdateStrategy))
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	return nil
}

// Convert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec(in *ManagedSeedSetSpec, out *seedmanagement.ManagedSeedSetSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedSetSpec_To_seedmanagement_ManagedSeedSetSpec(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec(in *seedmanagement.ManagedSeedSetSpec, out *ManagedSeedSetSpec, s conversion.Scope) error {
	out.Replicas = (*int32)(unsafe.Pointer(in.Replicas))
	out.Selector = in.Selector
	if err := Convert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate(&in.Template, &out.Template, s); err != nil {
		return err
	}
	if err := Convert_core_ShootTemplate_To_v1beta1_ShootTemplate(&in.ShootTemplate, &out.ShootTemplate, s); err != nil {
		return err
	}
	out.UpdateStrategy = (*UpdateStrategy)(unsafe.Pointer(in.UpdateStrategy))
	out.RevisionHistoryLimit = (*int32)(unsafe.Pointer(in.RevisionHistoryLimit))
	return nil
}

// Convert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec(in *seedmanagement.ManagedSeedSetSpec, out *ManagedSeedSetSpec, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedSetSpec_To_v1alpha1_ManagedSeedSetSpec(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus(in *ManagedSeedSetStatus, out *seedmanagement.ManagedSeedSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.NextReplicaNumber = in.NextReplicaNumber
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
	out.Conditions = *(*[]core.Condition)(unsafe.Pointer(&in.Conditions))
	out.PendingReplica = (*seedmanagement.PendingReplica)(unsafe.Pointer(in.PendingReplica))
	return nil
}

// Convert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus(in *ManagedSeedSetStatus, out *seedmanagement.ManagedSeedSetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedSetStatus_To_seedmanagement_ManagedSeedSetStatus(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus(in *seedmanagement.ManagedSeedSetStatus, out *ManagedSeedSetStatus, s conversion.Scope) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.ReadyReplicas = in.ReadyReplicas
	out.NextReplicaNumber = in.NextReplicaNumber
	out.CurrentReplicas = in.CurrentReplicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.CurrentRevision = in.CurrentRevision
	out.UpdateRevision = in.UpdateRevision
	out.CollisionCount = (*int32)(unsafe.Pointer(in.CollisionCount))
	out.Conditions = *(*[]v1beta1.Condition)(unsafe.Pointer(&in.Conditions))
	out.PendingReplica = (*PendingReplica)(unsafe.Pointer(in.PendingReplica))
	return nil
}

// Convert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus(in *seedmanagement.ManagedSeedSetStatus, out *ManagedSeedSetStatus, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedSetStatus_To_v1alpha1_ManagedSeedSetStatus(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(in *ManagedSeedSpec, out *seedmanagement.ManagedSeedSpec, s conversion.Scope) error {
	out.Shoot = (*seedmanagement.Shoot)(unsafe.Pointer(in.Shoot))
	if in.Gardenlet != nil {
		in, out := &in.Gardenlet, &out.Gardenlet
		*out = new(seedmanagement.Gardenlet)
		if err := Convert_v1alpha1_Gardenlet_To_seedmanagement_Gardenlet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Gardenlet = nil
	}
	return nil
}

// Convert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(in *ManagedSeedSpec, out *seedmanagement.ManagedSeedSpec, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(in *seedmanagement.ManagedSeedSpec, out *ManagedSeedSpec, s conversion.Scope) error {
	out.Shoot = (*Shoot)(unsafe.Pointer(in.Shoot))
	if in.Gardenlet != nil {
		in, out := &in.Gardenlet, &out.Gardenlet
		*out = new(Gardenlet)
		if err := Convert_seedmanagement_Gardenlet_To_v1alpha1_Gardenlet(*in, *out, s); err != nil {
			return err
		}
	} else {
		out.Gardenlet = nil
	}
	return nil
}

// Convert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(in *seedmanagement.ManagedSeedSpec, out *ManagedSeedSpec, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus(in *ManagedSeedStatus, out *seedmanagement.ManagedSeedStatus, s conversion.Scope) error {
	out.Conditions = *(*[]core.Condition)(unsafe.Pointer(&in.Conditions))
	out.ObservedGeneration = in.ObservedGeneration
	return nil
}

// Convert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus(in *ManagedSeedStatus, out *seedmanagement.ManagedSeedStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedStatus_To_seedmanagement_ManagedSeedStatus(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus(in *seedmanagement.ManagedSeedStatus, out *ManagedSeedStatus, s conversion.Scope) error {
	out.Conditions = *(*[]v1beta1.Condition)(unsafe.Pointer(&in.Conditions))
	out.ObservedGeneration = in.ObservedGeneration
	return nil
}

// Convert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus(in *seedmanagement.ManagedSeedStatus, out *ManagedSeedStatus, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedStatus_To_v1alpha1_ManagedSeedStatus(in, out, s)
}

func autoConvert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate(in *ManagedSeedTemplate, out *seedmanagement.ManagedSeedTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1alpha1_ManagedSeedSpec_To_seedmanagement_ManagedSeedSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate is an autogenerated conversion function.
func Convert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate(in *ManagedSeedTemplate, out *seedmanagement.ManagedSeedTemplate, s conversion.Scope) error {
	return autoConvert_v1alpha1_ManagedSeedTemplate_To_seedmanagement_ManagedSeedTemplate(in, out, s)
}

func autoConvert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate(in *seedmanagement.ManagedSeedTemplate, out *ManagedSeedTemplate, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_seedmanagement_ManagedSeedSpec_To_v1alpha1_ManagedSeedSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate is an autogenerated conversion function.
func Convert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate(in *seedmanagement.ManagedSeedTemplate, out *ManagedSeedTemplate, s conversion.Scope) error {
	return autoConvert_seedmanagement_ManagedSeedTemplate_To_v1alpha1_ManagedSeedTemplate(in, out, s)
}

func autoConvert_v1alpha1_PendingReplica_To_seedmanagement_PendingReplica(in *PendingReplica, out *seedmanagement.PendingReplica, s conversion.Scope) error {
	out.Name = in.Name
	out.Reason = seedmanagement.PendingReplicaReason(in.Reason)
	out.Since = in.Since
	out.Retries = (*int32)(unsafe.Pointer(in.Retries))
	return nil
}

// Convert_v1alpha1_PendingReplica_To_seedmanagement_PendingReplica is an autogenerated conversion function.
func Convert_v1alpha1_PendingReplica_To_seedmanagement_PendingReplica(in *PendingReplica, out *seedmanagement.PendingReplica, s conversion.Scope) error {
	return autoConvert_v1alpha1_PendingReplica_To_seedmanagement_PendingReplica(in, out, s)
}

func autoConvert_seedmanagement_PendingReplica_To_v1alpha1_PendingReplica(in *seedmanagement.PendingReplica, out *PendingReplica, s conversion.Scope) error {
	out.Name = in.Name
	out.Reason = PendingReplicaReason(in.Reason)
	out.Since = in.Since
	out.Retries = (*int32)(unsafe.Pointer(in.Retries))
	return nil
}

// Convert_seedmanagement_PendingReplica_To_v1alpha1_PendingReplica is an autogenerated conversion function.
func Convert_seedmanagement_PendingReplica_To_v1alpha1_PendingReplica(in *seedmanagement.PendingReplica, out *PendingReplica, s conversion.Scope) error {
	return autoConvert_seedmanagement_PendingReplica_To_v1alpha1_PendingReplica(in, out, s)
}

func autoConvert_v1alpha1_RollingUpdateStrategy_To_seedmanagement_RollingUpdateStrategy(in *RollingUpdateStrategy, out *seedmanagement.RollingUpdateStrategy, s conversion.Scope) error {
	out.Partition = (*int32)(unsafe.Pointer(in.Partition))
	return nil
}

// Convert_v1alpha1_RollingUpdateStrategy_To_seedmanagement_RollingUpdateStrategy is an autogenerated conversion function.
func Convert_v1alpha1_RollingUpdateStrategy_To_seedmanagement_RollingUpdateStrategy(in *RollingUpdateStrategy, out *seedmanagement.RollingUpdateStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_RollingUpdateStrategy_To_seedmanagement_RollingUpdateStrategy(in, out, s)
}

func autoConvert_seedmanagement_RollingUpdateStrategy_To_v1alpha1_RollingUpdateStrategy(in *seedmanagement.RollingUpdateStrategy, out *RollingUpdateStrategy, s conversion.Scope) error {
	out.Partition = (*int32)(unsafe.Pointer(in.Partition))
	return nil
}

// Convert_seedmanagement_RollingUpdateStrategy_To_v1alpha1_RollingUpdateStrategy is an autogenerated conversion function.
func Convert_seedmanagement_RollingUpdateStrategy_To_v1alpha1_RollingUpdateStrategy(in *seedmanagement.RollingUpdateStrategy, out *RollingUpdateStrategy, s conversion.Scope) error {
	return autoConvert_seedmanagement_RollingUpdateStrategy_To_v1alpha1_RollingUpdateStrategy(in, out, s)
}

func autoConvert_v1alpha1_Shoot_To_seedmanagement_Shoot(in *Shoot, out *seedmanagement.Shoot, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_Shoot_To_seedmanagement_Shoot is an autogenerated conversion function.
func Convert_v1alpha1_Shoot_To_seedmanagement_Shoot(in *Shoot, out *seedmanagement.Shoot, s conversion.Scope) error {
	return autoConvert_v1alpha1_Shoot_To_seedmanagement_Shoot(in, out, s)
}

func autoConvert_seedmanagement_Shoot_To_v1alpha1_Shoot(in *seedmanagement.Shoot, out *Shoot, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_seedmanagement_Shoot_To_v1alpha1_Shoot is an autogenerated conversion function.
func Convert_seedmanagement_Shoot_To_v1alpha1_Shoot(in *seedmanagement.Shoot, out *Shoot, s conversion.Scope) error {
	return autoConvert_seedmanagement_Shoot_To_v1alpha1_Shoot(in, out, s)
}

func autoConvert_v1alpha1_UpdateStrategy_To_seedmanagement_UpdateStrategy(in *UpdateStrategy, out *seedmanagement.UpdateStrategy, s conversion.Scope) error {
	out.Type = (*seedmanagement.UpdateStrategyType)(unsafe.Pointer(in.Type))
	out.RollingUpdate = (*seedmanagement.RollingUpdateStrategy)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_v1alpha1_UpdateStrategy_To_seedmanagement_UpdateStrategy is an autogenerated conversion function.
func Convert_v1alpha1_UpdateStrategy_To_seedmanagement_UpdateStrategy(in *UpdateStrategy, out *seedmanagement.UpdateStrategy, s conversion.Scope) error {
	return autoConvert_v1alpha1_UpdateStrategy_To_seedmanagement_UpdateStrategy(in, out, s)
}

func autoConvert_seedmanagement_UpdateStrategy_To_v1alpha1_UpdateStrategy(in *seedmanagement.UpdateStrategy, out *UpdateStrategy, s conversion.Scope) error {
	out.Type = (*UpdateStrategyType)(unsafe.Pointer(in.Type))
	out.RollingUpdate = (*RollingUpdateStrategy)(unsafe.Pointer(in.RollingUpdate))
	return nil
}

// Convert_seedmanagement_UpdateStrategy_To_v1alpha1_UpdateStrategy is an autogenerated conversion function.
func Convert_seedmanagement_UpdateStrategy_To_v1alpha1_UpdateStrategy(in *seedmanagement.UpdateStrategy, out *UpdateStrategy, s conversion.Scope) error {
	return autoConvert_seedmanagement_UpdateStrategy_To_v1alpha1_UpdateStrategy(in, out, s)
}
