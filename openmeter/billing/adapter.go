package billing

import (
	"context"
	"fmt"
)

type TxAdapter interface {
	ProfileAdapter

	Commit() error
	Rollback() error
}

type Adapter interface {
	ProfileAdapter

	WithTx(context.Context) (TxAdapter, error)
}

type ProfileAdapter interface {
	CreateProfile(ctx context.Context, input CreateProfileInput) (*Profile, error)
	GetProfile(ctx context.Context, input GetProfileInput) (*Profile, error)
	GetDefaultProfile(ctx context.Context, input GetDefaultProfileInput) (*Profile, error)
	DeleteProfile(ctx context.Context, input DeleteProfileInput) error
	UpdateProfile(ctx context.Context, input UpdateProfileAdapterInput) (*Profile, error)
}

type UpdateProfileAdapterInput struct {
	TargetState      Profile
	WorkflowConfigID string
}

func (i UpdateProfileAdapterInput) Validate() error {
	if err := i.TargetState.Validate(); err != nil {
		return fmt.Errorf("error validating target state profile: %w", err)
	}

	if i.TargetState.ID == "" {
		return fmt.Errorf("id is required")
	}

	if i.TargetState.UpdatedAt.IsZero() {
		return fmt.Errorf("updated at is required")
	}

	if i.WorkflowConfigID == "" {
		return fmt.Errorf("workflow config id is required")
	}

	return nil
}

func WithTxNoValue(ctx context.Context, repo Adapter, fn func(ctx context.Context, repo TxAdapter) error) error {
	var err error

	wrapped := func(ctx context.Context, repo TxAdapter) (interface{}, error) {
		if err = fn(ctx, repo); err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = WithTx(ctx, repo, wrapped)

	return err
}

func WithTx[T any](ctx context.Context, repo Adapter, fn func(ctx context.Context, repo TxAdapter) (T, error)) (resp T, err error) {
	var txRepo TxAdapter

	txRepo, err = repo.WithTx(ctx)
	if err != nil {
		return resp, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered from panic: %v: %w", r, err)

			if e := txRepo.Rollback(); e != nil {
				err = fmt.Errorf("failed to rollback transaction: %w: %w", e, err)
			}

			return
		}

		if err != nil {
			if e := txRepo.Rollback(); e != nil {
				err = fmt.Errorf("failed to rollback transaction: %w: %w", e, err)
			}

			return
		}

		if e := txRepo.Commit(); e != nil {
			err = fmt.Errorf("failed to commit transaction: %w", e)
		}
	}()

	resp, err = fn(ctx, txRepo)
	if err != nil {
		err = fmt.Errorf("failed to execute transaction: %w", err)
		return
	}

	return
}
