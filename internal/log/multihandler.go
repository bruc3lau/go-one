// file: internal/log/multihandler.go

package log

import (
	"context"
	"log/slog"
)

// MultiHandler 是一个 slog.Handler，它将日志记录分发到多个子 Handler。
type MultiHandler struct {
	handlers []slog.Handler
}

// NewMultiHandler 创建一个新的 MultiHandler。
func NewMultiHandler(handlers ...slog.Handler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

// Enabled 检查是否至少有一个子 Handler 启用了给定的级别。
func (h *MultiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	// 如果任何一个子 handler 启用了这个级别，那么多路复用 handler 就应该启用它
	for _, handler := range h.handlers {
		if handler.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

// Handle 将日志记录分发给所有的子 Handler。
func (h *MultiHandler) Handle(ctx context.Context, r slog.Record) error {
	// 将记录传递给每一个子 handler
	for _, handler := range h.handlers {
		// 注意：在一个循环中，我们通常不处理错误，因为我们希望所有 handler 都尝试处理日志。
		// 在更复杂的场景中，你可能需要收集和处理这些错误。
		_ = handler.Handle(ctx, r)
	}
	return nil
}

// WithAttrs 为所有的子 Handler 添加属性。
func (h *MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	// 创建一个新的 MultiHandler，其中包含已经更新了属性的子 handler
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithAttrs(attrs)
	}
	return &MultiHandler{handlers: newHandlers}
}

// WithGroup 为所有的子 Handler 添加分组。
func (h *MultiHandler) WithGroup(name string) slog.Handler {
	newHandlers := make([]slog.Handler, len(h.handlers))
	for i, handler := range h.handlers {
		newHandlers[i] = handler.WithGroup(name)
	}
	return &MultiHandler{handlers: newHandlers}
}
