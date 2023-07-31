package managedsecret

import (
	"context"
	"fmt"
	"smanager/config/repository"
	"smanager/internal/user"
)

type ManagedSecretService struct {
	ManagedSecretRepo repository.GenericRepository[ManagedSecret]
}

func NewManagedSercertService(ManagedSecretRepo repository.GenericRepository[ManagedSecret]) *ManagedSecretService {
	return &ManagedSecretService{ManagedSecretRepo}
}

func (msService *ManagedSecretService) CreateManagedSecret(ctx context.Context, secretName, secretValue string) error {
	resultQuery := make([]ManagedSecret, 0)
	userId := ctx.Value("user").(user.UserDTO).Id
	qbuilder := repository.QueriBuilder().With("name", secretName).With("user_id", userId)
	msService.ManagedSecretRepo.FindByParams(ctx, &resultQuery, qbuilder.Build())
	if len(resultQuery) != 0 {
		return fmt.Errorf("secret name already in use")
	}
	return msService.ManagedSecretRepo.Save(ctx, &ManagedSecret{
		Name:   secretName,
		Value:  secretValue,
		UserId: uint(userId),
	})
}

func (msService *ManagedSecretService) ListManagedSecret(ctx context.Context) ([]ManagedSecretDto, error) {
	secretList := make([]ManagedSecretDto, 0)

	resultQuery := make([]ManagedSecret, 0)
	userId := ctx.Value("user").(user.UserDTO).Id
	qbuilder := repository.QueriBuilder().With("user_id", userId)
	err := msService.ManagedSecretRepo.FindByParams(ctx, &resultQuery, qbuilder.Build())
	if err != nil {
		return secretList, fmt.Errorf("ListManagedSecret:%w", err)
	}
	for _, secret := range resultQuery {
		secretList = append(secretList, ManagedSecretDto{
			int(secret.ID),
			secret.Name,
		})
	}
	return secretList, nil
}

func (msService *ManagedSecretService) GetSecret(ctx context.Context, name string) (string, error) {
	userId := ctx.Value("user").(user.UserDTO).Id
	qbuilder := repository.QueriBuilder().With("user_id", userId).With("name", name)
	resultQuery := make([]ManagedSecret, 0)
	err := msService.ManagedSecretRepo.FindByParams(ctx, &resultQuery, qbuilder.Build())
	if len(resultQuery) > 0 {
		return resultQuery[0].Value, nil
	}
	return "", fmt.Errorf("GetSecret: %w", err)
}
func (msService *ManagedSecretService) EditManagedSecret() error {

	return nil
}

type ManagedSecretDto struct {
	ID   int
	Name string
}
