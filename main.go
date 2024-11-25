package main

import (
	"image"
	"image/color"
	"strings"

	"github.com/fogleman/gg"
)

// Fonts path
const fontRegular = "./assets/arial_regular.ttf"
const fontBold = "./assets/arial_bold.ttf"
const fontItalic = "./assets/arial_italic.ttf"
const fontBlack = "./assets/arial_black.ttf"

// Scale
const s = 2

// Design size
const w = 360 * s
const h = 2620 * s

const c = w / 2
const safeW = w - (32 * s)

func main() {
	// Current y position
	// This will increase as we draw
	var y float64 = 0

	dc := gg.NewContext(w, h)

	// Background
	dc.SetHexColor("#0D66C5")
	dc.Clear()
	dc.SetColor(color.White)

	imgLogo, _ := gg.LoadImage("./assets/logo_rekap.png")
	y += drawImageCenterFitWidth(dc, imgLogo, int(y), w)
	y -= 16 * s

	// ===============
	// Section 1
	// ===============
	dc.LoadFontFace(fontItalic, 10*s)
	y += drawString(dc, "Data per 1 Jan - 30 Nov 2024", c, y, safeW)
	y += 24 * s

	dc.SetHexColor("#FFB657")
	dc.LoadFontFace(fontBlack, 20*s)
	y += drawString(dc, "Hai, Rama Dewantara", c, y, safeW)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14*s)
	y += drawString(dc, "Tahun baru, lembaran baru! Raya siap jadi #KawanFinansial kamu setiap hari.", c, y, safeW)
	y += 24 * s

	y += drawString(dc, "Terima kasih, kamu sudah 3 tahun 5 Bulan bersama Raya! Selama barengan kamu adalah...", c, y, safeW)
	y += 16 * s

	imgTier, _ := gg.LoadImage("./assets/tiering.png")
	y += drawImageCenterFitWidth(dc, imgTier, int(y), w)

	dc.LoadFontFace(fontItalic, 14*s)
	y += drawString(dc, "“Kita belum terlalu banyak interaksi, cuma ketemu pas momen tertentu saja. Yuk, tahun depan lebih aktif nabung dan transaksi biar kita makin akrab!”", c, y, 275*s)
	y += 4 * s

	// ===============
	// Section 2
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0046A2")
	dc.Fill()

	imgSeparator1, _ := gg.LoadImage("./assets/sekat_1.png")
	y += drawImageCenterFitWidth(dc, imgSeparator1, int(y), w)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBlack, 16*s)
	y += drawString(dc, "Tahun 2024 ada banyak pencapaian yang kamu dapatkan!", c, y, safeW)
	y += 24 * s

	imgCashIn, _ := gg.LoadImage("./assets/uang_masuk.png")
	y += drawImageCenterFitWidth(dc, imgCashIn, int(y), 92*s)
	y += 8 * s

	dc.LoadFontFace(fontRegular, 13*s)
	y += drawString(dc, "Total Uang Masuk", c, y, safeW)
	y += 16 * s

	dc.LoadFontFace(fontBlack, 20*s)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8*s, "#35DDDA", 8*s)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14*s)
	y += drawString(dc, "Wah keren! Di tahun 2024 pemasukan terbesarmu ada di bulan April senilai", c, y, 275*s)
	y += 16 * s

	y += drawStringInBox(dc, "Rp.10.000.000", c, y, 4*s, "#A7FFFE", 4*s)
	y += 32 * s

	imgCashOut, _ := gg.LoadImage("./assets/uang_keluar.png")
	y += drawImageCenterFitWidth(dc, imgCashOut, int(y), 100*s)
	y += 8 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 13*s)
	y += drawString(dc, "Total Uang Keluar", c, y, safeW)
	y += 16 * s

	dc.LoadFontFace(fontBlack, 20*s)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8*s, "#FF8B00", 8*s)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14*s)
	y += drawString(dc, "Sedangkan pengeluaran terbesarmu tahun ini ada di bulan November senilai Rp200.000. Pastikan selalu atur keuanganmu pakai Saku Raya ya. ", c, y, 275*s)
	y += 32 * s

	dc.LoadFontFace(fontRegular, 13*s)
	y += drawString(dc, "Rata-rata saldo", c, y, safeW)
	y += 8 * s

	dc.LoadFontFace(fontBlack, 20*s)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8*s, "#FFC4E8", 8*s)
	y += 4 * s

	// ===============
	// Section 3
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0D66C5")
	dc.Fill()

	imgSeparator2, _ := gg.LoadImage("./assets/sekat_2.png")
	y += drawImageCenterFitWidth(dc, imgSeparator2, int(y), w)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBlack, 20*s)
	y += drawString(dc, "Transaksi kamu mencapai 350 kali!", c, y, 275*s)
	y += 16 * s

	dc.LoadFontFace(fontRegular, 12*s)
	y += drawString(dc, "Raya selalu siap untuk penuhi berbagai transaksimu.", c, y, safeW)
	y += 32 * s

	y += drawBarChart(dc, "#40D792", "1", "QRIS", "200", 0, y, 263*s)
	y += 16 * s
	y += drawBarChart(dc, "#FF6ED1", "2", "TopUp e-Wallet", "30", 0, y, 232*s)
	y += 16 * s
	y += drawBarChart(dc, "#FFB000", "3", "Transfer", "20", 0, y, 189*s)
	y += 16 * s
	y += drawBarChart(dc, "#24C8FF", "4", "Pulsa", "20", 0, y, 152*s)
	y += 16 * s
	y += drawBarChart(dc, "#7F8AF3", "5", "Lainnya", "20", 0, y, 115*s)
	y += 4 * s

	// ===============
	// Section 4
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0046A2")
	dc.Fill()

	imgSeparator3, _ := gg.LoadImage("./assets/sekat_3.png")
	y += drawImageCenterFitWidth(dc, imgSeparator3, int(y), w)
	y += 16 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBold, 16*s)
	y += drawString(dc, "Paling menyenangkan, selama pakai Raya total keuntungan kamu mencapai", c, y, safeW)
	y += 16 * s

	dc.LoadFontFace(fontBold, 20*s)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8*s, "#FFC4E8", 8*s)
	y += 8 * s

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 12*s)
	y += drawString(dc, "*Cashback dan Bunga", c, y, safeW)
	y += 4 * s

	// ===============
	// Section 5
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0D66C5")
	dc.Fill()

	imgSeparator4, _ := gg.LoadImage("./assets/sekat_4.png")
	y += drawImageCenterFitWidth(dc, imgSeparator4, int(y), w)

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBold, 16*s)
	y += drawString(dc, "Yuk, tingkatkan lagi hubungan ini dengan rutin nabung dan transaksi di Raya biar makin dekat ke depannya!", c, y, 275*s)
	y += 32 * s

	imgFooter, _ := gg.LoadImage("./assets/footer.png")
	y += drawImageCenterFitWidth(dc, imgFooter, int(y), safeW)

	// Export
	dc.SavePNG("output.png")
}

// Draw wrapped string in center and returns its height
func drawString(dc *gg.Context, txt string, x float64, y float64, w float64) float64 {
	const lineSpacing = 1.6

	splitLines := dc.WordWrap(txt, w)
	_, result := dc.MeasureMultilineString(strings.Join(splitLines, "\n"), lineSpacing)

	dc.DrawStringWrapped(txt, x, y, 0.5, 0, w, lineSpacing, gg.AlignCenter)

	return result
}

// Draw string in a rounded rectangle and returns its height
func drawStringInBox(dc *gg.Context, txt string, x float64, y float64, r float64, hex string, padding float64) float64 {
	sw, sh := dc.MeasureString(txt)

	swp := sw + padding*2
	shp := sh + padding*2

	rX := x - (swp / 2)

	dc.DrawRoundedRectangle(rX, y, swp, shp, r)
	dc.SetHexColor(hex)
	dc.Fill()
	dc.SetColor(color.Black)
	dc.DrawRoundedRectangle(rX, y, swp, shp, r)
	dc.Stroke()

	dc.DrawStringAnchored(txt, x, y+(shp/2), 0.5, 0.5)

	return shp
}

// Draw bar with number, name and count. Returns fixed bar height
func drawBarChart(dc *gg.Context, hex string, number string, name string, count string, x float64, y float64, w float64) float64 {
	const barH = 35.0 * s

	dc.DrawRectangle(x, y, w, barH)
	dc.SetHexColor(hex)
	dc.Fill()
	dc.SetColor(color.Black)
	dc.DrawRectangle(x, y, w, barH)
	dc.Stroke()

	cx := x + w - barH/2 - (4 * s)
	cy := y + barH/2
	cr := barH/2 - (4 * s)

	dc.DrawCircle(cx, cy, cr)
	dc.Stroke()

	dc.LoadFontFace(fontBold, 18*s)
	dc.DrawStringAnchored(number, cx, cy, 0.5, 0.5)

	nx := x + w + (16 * s)
	ny := y
	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 9*s)
	dc.DrawStringAnchored(name, nx, ny, 0, 1)
	ny += dc.FontHeight()
	ny += 8 * s

	dc.SetHexColor(hex)
	dc.LoadFontFace(fontBlack, 26*s)
	dc.DrawStringAnchored(count, nx, ny, 0, 1)

	cnw, _ := dc.MeasureString(count)
	nx += cnw + (4 * s)
	ny += dc.FontHeight() / 2

	dc.LoadFontFace(fontBold, 13*s)
	dc.DrawStringAnchored("x", nx, ny, 0, 0.5)

	return barH
}

// Draw image at center, fit to w and returns its scaled height
func drawImageCenterFitWidth(dc *gg.Context, img image.Image, y, w int) float64 {
	imw := img.Bounds().Dx()
	imh := img.Bounds().Dy()

	sw := float64(w) / float64(imw)
	simh := sw * float64(imh)

	imc := gg.NewContext(w, int(simh))
	imc.Scale(float64(sw), float64(sw))
	imc.DrawImage(img, 0, 0)

	dc.DrawImageAnchored(imc.Image(), c, y, 0.5, 0)

	return float64(simh)
}
