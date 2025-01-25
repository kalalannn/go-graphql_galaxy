package server

import (
	"context"
	"go-graphql_galaxy/internal/gqlcontext"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/errcode"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

const (
	depthExtensionName = "DepthExtension"
	errDepthLimit      = "DEPTH_LIMIT_EXCEEDED"
)

const (
	IntrospectionQueryName    = "IntrospectionQuery"
	IntrospectionObjectSchema = "__schema"
)

const (
	CharacterObjectSchema  = "character"
	CharactersObjectSchema = "characters"
	NemesisObjectSchema    = "nemesis"
	NemesesObjectSchema    = "nemeses"
	SecretObjectSchema     = "secret"
	SecretsObjectSchema    = "secrets"
)

var _ interface {
	graphql.OperationContextMutator
	graphql.HandlerExtension
} = &DepthExtension{}

type DepthExtension struct {
	MaxDepth int
}

func NewDepthExtension(limit int) *DepthExtension {
	return &DepthExtension{
		MaxDepth: limit,
	}
}

func (m DepthExtension) ExtensionName() string {
	return depthExtensionName
}

func (m *DepthExtension) Validate(schema graphql.ExecutableSchema) error {
	return nil
}

func (m *DepthExtension) MutateOperationContext(ctx context.Context, opCtx *graphql.OperationContext) *gqlerror.Error {
	if opCtx.Operation.Name == IntrospectionQueryName {
		return nil
	}

	if gqlcontext.DFSMaxDepth(opCtx, m.MaxDepth) != nil {
		err := gqlerror.Errorf("Max depth limit exceeded > %d", m.MaxDepth)
		errcode.Set(err, errDepthLimit)
		return err
	}
	return nil
}

func (m *DepthExtension) InterceptRootField(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
	rootFieldCtx := graphql.GetRootFieldContext(ctx)

	switch rootFieldCtx.Object {
	case CharacterObjectSchema, CharactersObjectSchema,
		NemesisObjectSchema, NemesesObjectSchema,
		SecretObjectSchema, SecretsObjectSchema:
		return next(context.WithValue(ctx,
			gqlcontext.PreloadContextKey,
			gqlcontext.DFSPreload(rootFieldCtx)))
	}

	return next(ctx)
}

// func (m *DepthExtension) InterceptOperation(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
// 	log.Debug("InterceptOperation: before")
// 	return next(ctx)
// }

// func (m *DepthExtension) InterceptField(ctx context.Context, next graphql.Resolver) (res any, err error) {
// 	log.Debug("InterceptField: before")
// 	return next(ctx)
// }
