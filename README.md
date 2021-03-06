# Docs

## [faker](#)
Faker  will generate you a fake data based on your Struct.


** This is a fork of https://github.com/bxcodec/faker which was customized to work for Slug **

The differences between this version and bxcodec are:
* bxcodec relies on a function to generate fake data with settings configured globally
* this version you construct an instance of the fake data generator and can specify configuration on that instance (this was needed to scope the fake data configuration to a test as it can very between tests)
* bxcodec relies on tags defined on the struct to determine fake-data type
* this version let's you define this configuration by field name

In general you can use the same details as noted in the docs [Godoc](https://godoc.org/github.com/bxcodec/faker). The main difference
is instead of using 

```
faker.FakeData(<your instance>)
```
you
```
   fakeGen := faker.NewFakeGenerator()
    ... additional config
   fakeGen.FakeData(<your instance>)
```

this additional configuration is:

* you can specify a regex to ignore certain fields. This is done via the method AddFieldFilter giving it a regex to match field names to exclude from filling
* you can specify a tag on a field by name. This is done by via the method AddFieldTag giving it the field name and the tag
* you can specify additional value providers. This is really used to assign a specific value to a field by name where you specific the field name and give a provider used to get the value for that field.

## Index

* [Support](#support)
* [Getting Started](#getting-started)
* [Example](#example)
* [Limitation](#limitation)
* [Contribution](#contribution)


## Support

You can file an [Issue](https://github.com/bxcodec/faker/issues/new).
See documentation in [Godoc](https://godoc.org/github.com/bxcodec/faker)


## Getting Started

#### Download

```shell
go get -u github.com/bxcodec/faker/v3
```
# Example

---
 
 - Using Struct's tag:  [WithStructTag.md](/WithStructTag.md)
 - Custom Struct's tag (define your own faker data): [CustomFaker.md](/CustomFaker.md)
 - Without struct's tag: [WithoutTag.md](/WithoutTag.md)
 - Single Fake Data Function: [SingleFakeData.md](/SingleFakeData.md)
 
## DEMO

---

![Example to use Faker](https://cdn-images-1.medium.com/max/800/1*AkMbxngg7zfvtWiuvFb4Mg.gif)

## Benchmark

---

Bench To Generate Fake Data
#### Without Tag
```bash
BenchmarkFakerDataNOTTagged-4             500000              3049 ns/op             488 B/op         20 allocs/op
```

#### Using Tag
```bash
 BenchmarkFakerDataTagged-4                100000             17470 ns/op             380 B/op         26 allocs/op
```

### MUST KNOW

---

The Struct Field must PUBLIC.<br>
Support Only For :

* int  int8  int16  int32  int64
* []int  []int8  []int16  []int32  []int64
* bool []bool
* string []string
* float32 float64 []float32 []float64
* Nested Struct Field
* time.Time []time.Time

## Limitation

---

Unfortunately this library has some limitation
* It does not support private fields. Make sure your structs fields you intend to generate fake data for are public, it would otherwise trigger a panic. You can however omit fields using a tag skip `faker:"-"` on your private fields.
* It does not support the `interface{}` data type. How could we generate anything without knowing its data type?
* It does not support the `map[interface{}]interface{}, map[any_type]interface{}, map[interface{}]any_type` data types. Once again, we cannot generate values for an unknown data type.
* Custom types are not fully supported. However some custom types are already supported: we are still investigating how to do this the correct way. For now, if you use `faker`, it's safer not to use any custom types in order to avoid panics.

## Contribution

---

To contrib to this project, you can open a PR or an issue.
