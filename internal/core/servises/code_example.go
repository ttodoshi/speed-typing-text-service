package servises

import (
	"github.com/jinzhu/copier"
	"speed-typing-text-service/internal/adapters/dto"
	"speed-typing-text-service/internal/core/domain"
	"speed-typing-text-service/internal/core/errors"
	"speed-typing-text-service/internal/core/ports"
	"speed-typing-text-service/pkg/logging"
)

type CodeExampleService struct {
	repo ports.CodeExampleRepository
	log  logging.Logger
}

func NewCodeExampleService(repo ports.CodeExampleRepository, log logging.Logger) ports.CodeExampleService {
	return &CodeExampleService{
		repo: repo,
		log:  log,
	}
}

func (s *CodeExampleService) GetProgrammingLanguages() (getProgrammingLanguagesDto []dto.GetProgrammingLanguageDto, err error) {
	programmingLanguages := s.repo.GetProgrammingLanguages()

	err = copier.Copy(&getProgrammingLanguagesDto, &programmingLanguages)
	if err != nil {
		err = &errors.MappingError{Message: `struct mapping error`}
	}
	return
}

func (s *CodeExampleService) GetCodeExampleByUUID(UUID string) (getCodeExampleDto dto.GetCodeExampleDto, err error) {
	var codeExample domain.CodeExample
	codeExample, err = s.repo.GetCodeExampleByUUID(UUID)
	if err != nil {
		return
	}

	err = copier.Copy(&getCodeExampleDto, &codeExample)
	if err != nil {
		err = &errors.MappingError{Message: `struct mapping error`}
	}
	return
}

func (s *CodeExampleService) GetCodeExamples() (getCodeExamplesDto []dto.GetCodeExampleDto, err error) {
	codeExamples := s.repo.GetCodeExamples()

	err = copier.Copy(&getCodeExamplesDto, &codeExamples)
	if err != nil {
		err = &errors.MappingError{Message: `struct mapping error`}
	}
	return
}

func (s *CodeExampleService) GetCodeExamplesByProgrammingLanguageName(programmingLanguageName string) (getCodeExamplesDto []dto.GetCodeExampleDto, err error) {
	codeExamples, err := s.repo.GetCodeExamplesByProgrammingLanguageName(programmingLanguageName)
	if err != nil {
		s.log.Infof(`error getting code examples by programming language '%s' due to error: %v`, programmingLanguageName, err)
		return
	}

	err = copier.Copy(&getCodeExamplesDto, &codeExamples)
	if err != nil {
		err = &errors.MappingError{Message: `struct mapping error`}
	}
	return
}
