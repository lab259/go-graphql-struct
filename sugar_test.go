package gqlstruct_test

import (
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sugar", func() {
	Describe("Description", func() {
		It("should apply the description to a field", func() {
			field := graphql.Field{}
			err := gqlstruct.WithDescription("Description 1").Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.Description).To(Equal("Description 1"))
		})

		It("should apply the description to an argument", func() {
			argument := graphql.ArgumentConfig{}
			err := gqlstruct.WithDescription("Description 1").Apply(&argument)
			Expect(err).ToNot(HaveOccurred())
			Expect(argument.Description).To(Equal("Description 1"))
		})

		It("should apply the description to an object", func() {
			obj := graphql.ObjectConfig{}
			err := gqlstruct.WithDescription("Description 1").Apply(&obj)
			Expect(err).ToNot(HaveOccurred())
			Expect(obj.Description).To(Equal("Description 1"))
		})

		It("should fail applying the description to a not supported object", func() {
			obj := map[string]interface{}{}
			err := gqlstruct.WithDescription("Description 1").Apply(&obj)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not supported"))
			Expect(err.Error()).To(ContainSubstring("map"))
		})
	})

	Describe("DefaultValue", func() {
		It("should apply the default value to an argument", func() {
			argument := graphql.ArgumentConfig{}
			err := gqlstruct.WithDefaultvalue("default value").Apply(&argument)
			Expect(err).ToNot(HaveOccurred())
			Expect(argument.DefaultValue).To(Equal("default value"))
		})

		It("should fail applying the default value to a not supported object", func() {
			field := graphql.Field{}
			err := gqlstruct.WithDefaultvalue("default value").Apply(&field)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not supported"))
			Expect(err.Error()).To(ContainSubstring("Field"))
		})
	})

	Describe("DeprecationReason", func() {
		It("should apply the deprecation reason to a field", func() {
			field := graphql.Field{}
			err := gqlstruct.WithDeprecationReason("deprecation reason").Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.DeprecationReason).To(Equal("deprecation reason"))
		})

		It("should fail applying the deprecation reason to a not supported object", func() {
			argument := graphql.ArgumentConfig{}
			err := gqlstruct.WithDeprecationReason("default value").Apply(&argument)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not supported"))
			Expect(err.Error()).To(ContainSubstring("ArgumentConfig"))
		})
	})

	Describe("Resolve", func() {
		It("should apply the resolver to a field", func() {
			field := graphql.Field{}
			err := gqlstruct.WithResolve(func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			}).Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.Resolve).NotTo(BeNil())
		})

		It("should fail applying the resolver to a not supported object", func() {
			argument := graphql.ArgumentConfig{}
			err := gqlstruct.WithResolve(func(p graphql.ResolveParams) (interface{}, error) {
				return nil, nil
			}).Apply(&argument)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not supported"))
			Expect(err.Error()).To(ContainSubstring("ArgumentConfig"))
		})
	})
})
