package usermodel

import (
	"os/exec"
	"strings"

	"github.com/xyproto/files"
)

type Task string

// Tasks
const (
	llmManagerExecutable = "llm-manager"

	ChatTask           = "chat"
	CodeCompletionTask = "code-completion"
	TestTask           = "test"
	TextGenerationTask = "text-generation"
	ToolUseTask        = "tool-use"
	TranslationTask    = "translation"
	VisionTask         = "vision"
)

var DefaultModel = "gemma2:2b"

func AvailableTasks() []Task {
	return []Task{ChatTask, CodeCompletionTask, TestTask, TextGenerationTask, ToolUseTask, TranslationTask, VisionTask}
}

func GetChatModel() string           { return Get(ChatTask) }
func GetCodeCompletionModel() string { return Get(CodeCompletionTask) }
func GetTestModel() string           { return Get(TestTask) }
func GetTextGenerationModel() string { return Get(TextGenerationTask) }
func GetToolUseModel() string        { return Get(ToolUseTask) }
func GetTranslationModel() string    { return Get(TranslationTask) }
func GetvisionModel() string         { return Get(VisionTask) }

// Get attempts to retrieve the model name using llm-manager.
// If llm-manager is not available or the command fails, it falls back to the DefaultModel variable.
func Get(task Task) string {
	llmManagerPath := files.WhichCached(llmManagerExecutable)
	if llmManagerPath == "" {
		return DefaultModel
	}
	cmd := exec.Command(llmManagerPath, "get", string(task))
	outputBytes, err := cmd.Output()
	if err != nil {
		return DefaultModel
	}
	output := strings.TrimSpace(string(outputBytes))
	if output == "" {
		return DefaultModel
	}
	return output
}
