package gqlstruct_test

import (
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("ArrayOf", func() {
	Describe("String", func() {
		It("should generate graphql.Type with a string array field", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(""))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[String]"))
		})

		It("should generate graphql.Type with a string pointer array field", func() {
			s := ""
			obj := gqlstruct.ArrayOf(reflect.TypeOf(&s))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[String]"))
		})
	})

	Describe("Int", func() {
		It("should generate graphql.Type with a int array field", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(123))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Int]"))
		})

		It("should generate graphql.Type with a int pointer array field", func() {
			s := 123
			obj := gqlstruct.ArrayOf(reflect.TypeOf(&s))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Int]"))
		})
	})

	Describe("Float", func() {
		It("should generate graphql.Type with a float array field", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(123.))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Float]"))
		})

		It("should generate graphql.Type with a float pointer array field", func() {
			s := 123.
			obj := gqlstruct.ArrayOf(reflect.TypeOf(&s))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Float]"))
		})
	})

	Describe("Custom Type", func() {
		It("should generate graphql.Type with a int array field", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(CustomFieldType{}))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})

		It("should generate graphql.Type with a int pointer array field", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(&CustomFieldType{}))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})
	})
})
