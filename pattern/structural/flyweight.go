package structural

// 享元模式（Flyweight）：运用共享技术有效地支持大量细粒度对象的复用，节省内存。
//
// 关键：区分“内部状态”（可共享、与场景无关）与“外部状态”（随场景变化、由调用方传入）。
// 典型场景：字符渲染、连接池、棋子等大量重复对象。

// Glyph 享元接口：内部状态由享元持有，外部状态通过参数传入。
type Glyph interface {
	Render(font string) string // font 为外部状态
}

// CharGlyph 具体享元：共享的字符字形，code 是内部状态。
type CharGlyph struct {
	code rune
}

func (g CharGlyph) Render(font string) string {
	return string(g.code) + "@" + font
}

// GlyphFactory 享元工厂：缓存已创建的享元，重复请求返回同一实例。
type GlyphFactory struct {
	pool map[rune]Glyph
}

func NewGlyphFactory() *GlyphFactory { return &GlyphFactory{pool: map[rune]Glyph{}} }

func (f *GlyphFactory) Get(code rune) Glyph {
	if g, ok := f.pool[code]; ok {
		return g
	}
	g := CharGlyph{code: code}
	f.pool[code] = g
	return g
}
