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
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(""))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[String]"))
		})

		It("should generate graphql.Type with a string pointer array field", func() {
			s := ""
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(&s))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[String]"))
		})
	})

	Describe("Int", func() {
		It("should generate graphql.Type with a int array field", func() {
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(123))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Int]"))
		})

		It("should generate graphql.Type with a int pointer array field", func() {
			s := 123
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(&s))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Int]"))
		})
	})

	Describe("Float", func() {
		It("should generate graphql.Type with a float array field", func() {
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(123.))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Float]"))
		})

		It("should generate graphql.Type with a float pointer array field", func() {
			s := 123.
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(&s))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[Float]"))
		})
	})

	Describe("Custom Type", func() {
		It("should generate graphql.Type with a int array field", func() {
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(CustomFieldType{}))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})

		It("should generate graphql.Type with a int pointer array field", func() {
			obj, err := gqlstruct.NewEncoder().ArrayOf(reflect.TypeOf(&CustomFieldType{}))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})

		It("should generate graphql.Type with a int pointer array field from the cache", func() {
			enc := gqlstruct.NewEncoder()

			obj, err := enc.ArrayOf(reflect.TypeOf(&CustomFieldType{}))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))

			obj, err = enc.ArrayOf(reflect.TypeOf(&CustomFieldType{}))
			Expect(err).ToNot(HaveOccurred())
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})

		It("should panic when find a type not supported", func() {
			type CustomFieldTypeWithNoGraphQLConverted interface{}

			type StructExampleNotFound struct {
				Field1 []CustomFieldTypeWithNoGraphQLConverted `graphql:"field1"`
			}

			Expect(func() {
				gqlstruct.ArrayOf(reflect.TypeOf(&StructExampleNotFound{}))
			}).To(Panic())
		})

		It("should generate graphql.Type with a int pointer array field using the global", func() {
			obj := gqlstruct.ArrayOf(reflect.TypeOf(&CustomFieldType{}))
			Expect(obj).ToNot(BeNil())
			Expect(obj.String()).To(Equal("[CustomFieldType]"))
		})
	})
})
