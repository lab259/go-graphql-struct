package gqlstruct_test

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type erroredOption struct{}

func (*erroredOption) Apply(dst interface{}) error {
	return errors.New("forced error")
}

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

	Describe("Args", func() {
		It("should apply the arguments to a field", func() {
			type Args struct {
				Name string `graphql:"name"`
				Age  int    `graphql:"age"`
			}

			field := graphql.Field{}
			err := gqlstruct.WithArgs(Args{}).Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.Args).To(HaveLen(2))
			Expect(field.Args).To(HaveKey("name"))
			Expect(field.Args).To(HaveKey("age"))
		})

		It("should apply the arguments to a field with a custom encoder", func() {
			enc := gqlstruct.NewEncoder()

			type Args struct {
				Name string `graphql:"name"`
				Age  int    `graphql:"age"`
			}

			field := graphql.Field{}
			err := gqlstruct.WithArgs(enc, Args{}).Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.Args).To(HaveLen(2))
			Expect(field.Args).To(HaveKey("name"))
			Expect(field.Args).To(HaveKey("age"))
		})

		It("should panic when the first args is not a encoder", func() {
			type Args struct {
				Name string `graphql:"name"`
				Age  int    `graphql:"age"`
			}

			field := graphql.Field{}
			Expect(func() {
				gqlstruct.WithArgs(Args{}, Args{}).Apply(&field)
			}).To(Panic())
		})

		It("should panic when the call with more than 2 arguments", func() {
			field := graphql.Field{}
			Expect(func() {
				gqlstruct.WithArgs(1, 2, 3).Apply(&field)
			}).To(Panic())
		})

		It("should fail due to the arg is not a struct", func() {
			field := graphql.Field{}
			err := gqlstruct.WithArgs([]interface{}{}).Apply(&field)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("cannot build args from a non struct"))
		})

		It("should fail due to not supported data type", func() {
			type Args struct {
				Field1 []interface{} `graphql:"field1"`
			}

			field := graphql.Field{}
			err := gqlstruct.WithArgs(Args{}).Apply(&field)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not recognized"))
			Expect(err.Error()).To(ContainSubstring("interface {}"))
		})

		It("should fail when not applying to a field", func() {
			type Args struct {
				Field1 []interface{} `graphql:"field1"`
			}

			obj := graphql.ObjectConfig{}
			err := gqlstruct.WithArgs(Args{}).Apply(&obj)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("not supported"))
			Expect(err.Error()).To(ContainSubstring("ObjectConfig"))
		})
	})

	Describe("Type", func() {
		type ThisIsAType struct {
			Name string `graphql:"name"`
		}
		t := gqlstruct.Struct(ThisIsAType{})

		It("should apply the type to a field", func() {
			field := graphql.Field{}
			err := gqlstruct.WithType(t).Apply(&field)
			Expect(err).ToNot(HaveOccurred())
			Expect(field.Type).To(Equal(t))
		})

		It("should fail applying the resolver to a not supported object", func() {
			argument := graphql.ArgumentConfig{}
			err := gqlstruct.WithType(t).Apply(&argument)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("is not supported"))
			Expect(err.Error()).To(ContainSubstring("ArgumentConfig"))
		})
	})
})
