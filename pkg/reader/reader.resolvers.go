package reader

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.41

import (
	"context"
	"log/slog"

	"github.com/calindra/cartesi-rollups-graphql/pkg/reader/graph"
	"github.com/calindra/cartesi-rollups-graphql/pkg/reader/model"
)

// Vouchers is the resolver for the vouchers field.
func (r *inputResolver) Vouchers(ctx context.Context, obj *model.Input, first *int, last *int, after *string, before *string) (*model.Connection[*model.Voucher], error) {
	if first == nil && last == nil && after == nil && before == nil {
		return r.adapter.GetAllVouchersByInputIndex(ctx, &obj.Index)
	}
	return r.adapter.GetVouchers(ctx, first, last, after, before, &obj.Index, nil)
}

// Notices is the resolver for the notices field.
func (r *inputResolver) Notices(ctx context.Context, obj *model.Input, first *int, last *int, after *string, before *string) (*model.Connection[*model.Notice], error) {
	if first == nil && last == nil && after == nil && before == nil {
		return r.adapter.GetAllNoticesByInputIndex(ctx, &obj.Index)
	}
	return r.adapter.GetNotices(ctx, first, last, after, before, &obj.Index)
}

// Reports is the resolver for the reports field.
func (r *inputResolver) Reports(ctx context.Context, obj *model.Input, first *int, last *int, after *string, before *string) (*model.Connection[*model.Report], error) {
	if first == nil && last == nil && after == nil && before == nil {
		return r.adapter.GetAllReportsByInputIndex(ctx, &obj.Index)
	}
	return r.adapter.GetReports(ctx, first, last, after, before, &obj.Index)
}

// Input is the resolver for the input field.
func (r *noticeResolver) Input(ctx context.Context, obj *model.Notice) (*model.Input, error) {
	slog.Debug("Find input by index", "inputIndex", obj.InputIndex)
	input, err := r.adapter.GetInputByIndex(ctx, obj.InputIndex)
	if err != nil {
		slog.Error("Input not found")
		return nil, err
	}
	return input, nil
}

// Input is the resolver for the input field.
func (r *queryResolver) Input(ctx context.Context, id string) (*model.Input, error) {
	slog.Debug("queryResolver.Input", "id", id)
	return r.adapter.GetInput(ctx, id)
}

// Voucher is the resolver for the voucher field.
func (r *queryResolver) Voucher(ctx context.Context, outputIndex int) (*model.Voucher, error) {
	return r.adapter.GetVoucher(ctx, outputIndex)
}

// Notice is the resolver for the notice field.
func (r *queryResolver) Notice(ctx context.Context, outputIndex int) (*model.Notice, error) {
	return r.adapter.GetNotice(ctx, outputIndex)
}

// Report is the resolver for the report field.
func (r *queryResolver) Report(ctx context.Context, reportIndex int) (*model.Report, error) {
	return r.adapter.GetReport(ctx, reportIndex)
}

// Inputs is the resolver for the inputs field.
func (r *queryResolver) Inputs(ctx context.Context, first *int, last *int, after *string, before *string, where *model.InputFilter) (*model.Connection[*model.Input], error) {
	return r.adapter.GetInputs(ctx, first, last, after, before, where)
}

// Vouchers is the resolver for the vouchers field.
func (r *queryResolver) Vouchers(ctx context.Context, first *int, last *int, after *string, before *string, filter []*model.ConvenientFilter) (*model.Connection[*model.Voucher], error) {
	return r.adapter.GetVouchers(ctx, first, last, after, before, nil, filter)
}

// Notices is the resolver for the notices field.
func (r *queryResolver) Notices(ctx context.Context, first *int, last *int, after *string, before *string) (*model.Connection[*model.Notice], error) {
	return r.adapter.GetNotices(ctx, first, last, after, before, nil)
}

// Reports is the resolver for the reports field.
func (r *queryResolver) Reports(ctx context.Context, first *int, last *int, after *string, before *string) (*model.Connection[*model.Report], error) {
	return r.adapter.GetReports(ctx, first, last, after, before, nil)
}

// Input is the resolver for the input field.
func (r *reportResolver) Input(ctx context.Context, obj *model.Report) (*model.Input, error) {
	return r.adapter.GetInputByIndex(ctx, obj.InputIndex)
}

// Input is the resolver for the input field.
func (r *voucherResolver) Input(ctx context.Context, obj *model.Voucher) (*model.Input, error) {
	return r.adapter.GetInputByIndex(ctx, obj.InputIndex)
}

// Input returns graph.InputResolver implementation.
func (r *Resolver) Input() graph.InputResolver { return &inputResolver{r} }

// Notice returns graph.NoticeResolver implementation.
func (r *Resolver) Notice() graph.NoticeResolver { return &noticeResolver{r} }

// Query returns graph.QueryResolver implementation.
func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

// Report returns graph.ReportResolver implementation.
func (r *Resolver) Report() graph.ReportResolver { return &reportResolver{r} }

// Voucher returns graph.VoucherResolver implementation.
func (r *Resolver) Voucher() graph.VoucherResolver { return &voucherResolver{r} }

type inputResolver struct{ *Resolver }
type noticeResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type reportResolver struct{ *Resolver }
type voucherResolver struct{ *Resolver }
