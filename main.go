package main

import (
	"github.com/fogleman/gg"
	"image/color"
	"strings"
)

const fontRegular = "./assets/arial_regular.ttf"
const fontBold = "./assets/arial_bold.ttf"
const fontItalic = "./assets/arial_italic.ttf"
const fontBlack = "./assets/arial_black.ttf"

func main() {
	const w = 360
	const h = 2620

	const c = w / 2
	const safeW = w - 32

	// Current y position
	// This will increase as we draw
	var y float64 = 0

	dc := gg.NewContext(w, h)

	// Background
	dc.SetHexColor("#0D66C5")
	dc.Clear()
	dc.SetColor(color.White)

	imgLogo, _ := gg.LoadImage("./assets/logo_rekap_360.png")
	dc.DrawImage(imgLogo, 0, 0)
	y += float64(imgLogo.Bounds().Dy()) - 16

	// ===============
	// Section 1
	// ===============
	dc.LoadFontFace(fontItalic, 10)
	y += drawString(dc, "Data per 1 Jan - 30 Nov 2024", c, y, safeW)
	y += 24

	dc.SetHexColor("#FFB657")
	dc.LoadFontFace(fontBlack, 20)
	y += drawString(dc, "Hai, Rama Dewantara", c, y, safeW)
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14)
	y += drawString(dc, "Tahun baru, lembaran baru! Raya siap jadi #KawanFinansial kamu setiap hari.", c, y, safeW)
	y += 24

	y += drawString(dc, "Terima kasih, kamu sudah 3 tahun 5 Bulan bersama Raya! Selama barengan kamu adalah...", c, y, safeW)
	y += 16

	imgTier, _ := gg.LoadImage("./assets/tiering_360.png")
	dc.DrawImage(imgTier, 0, int(y))
	y += float64(imgTier.Bounds().Dy())

	dc.LoadFontFace(fontItalic, 14)
	y += drawString(dc, "“Kita belum terlalu banyak interaksi, cuma ketemu pas momen tertentu saja. Yuk, tahun depan lebih aktif nabung dan transaksi biar kita makin akrab!”", c, y, 275)
	y += 4

	// ===============
	// Section 2
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0046A2")
	dc.Fill()

	imgSeparator1, _ := gg.LoadImage("./assets/sekat_1_360.png")
	dc.DrawImage(imgSeparator1, 0, int(y))
	y += float64(imgSeparator1.Bounds().Dy())
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBlack, 16)
	y += drawString(dc, "Tahun 2024 ada banyak pencapaian yang kamu dapatkan!", c, y, safeW)
	y += 24

	imgCashIn, _ := gg.LoadImage("./assets/uang_masuk.png")
	dc.DrawImageAnchored(imgCashIn, c, int(y), 0.5, 0)
	y += float64(imgCashIn.Bounds().Dy())
	y += 8

	dc.LoadFontFace(fontRegular, 13)
	y += drawString(dc, "Total Uang Masuk", c, y, safeW)
	y += 16

	dc.LoadFontFace(fontBlack, 20)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8, "#35DDDA", 8)
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14)
	y += drawString(dc, "Wah keren! Di tahun 2024 pemasukan terbesarmu ada di bulan April senilai", c, y, 275)
	y += 16

	y += drawStringInBox(dc, "Rp.10.000.000", c, y, 4, "#A7FFFE", 4)
	y += 32

	imgCashOut, _ := gg.LoadImage("./assets/uang_keluar.png")
	dc.DrawImageAnchored(imgCashOut, c, int(y), 0.5, 0)
	y += float64(imgCashOut.Bounds().Dy())
	y += 8

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 13)
	y += drawString(dc, "Total Uang Keluar", c, y, safeW)
	y += 16

	dc.LoadFontFace(fontBlack, 20)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8, "#FF8B00", 8)
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 14)
	y += drawString(dc, "Sedangkan pengeluaran terbesarmu tahun ini ada di bulan November senilai Rp200.000. Pastikan selalu atur keuanganmu pakai Saku Raya ya. ", c, y, 275)
	y += 32

	dc.LoadFontFace(fontRegular, 13)
	y += drawString(dc, "Rata-rata saldo", c, y, safeW)
	y += 8

	dc.LoadFontFace(fontBlack, 20)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8, "#FFC4E8", 8)
	y += 4

	// ===============
	// Section 3
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0D66C5")
	dc.Fill()

	imgSeparator2, _ := gg.LoadImage("./assets/sekat_2_360.png")
	dc.DrawImage(imgSeparator2, 0, int(y))
	y += float64(imgSeparator2.Bounds().Dy())
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBlack, 20)
	y += drawString(dc, "Transaksi kamu mencapai 350 kali!", c, y, 275)
	y += 16

	dc.LoadFontFace(fontRegular, 12)
	y += drawString(dc, "Raya selalu siap untuk penuhi berbagai transaksimu.", c, y, safeW)
	y += 32

	y += drawBarChart(dc, "#40D792", "1", "QRIS", "200", 0, y, 263)
	y += 16
	y += drawBarChart(dc, "#FF6ED1", "2", "TopUp e-Wallet", "30", 0, y, 232)
	y += 16
	y += drawBarChart(dc, "#FFB000", "3", "Transfer", "20", 0, y, 189)
	y += 16
	y += drawBarChart(dc, "#24C8FF", "4", "Pulsa", "20", 0, y, 152)
	y += 16
	y += drawBarChart(dc, "#7F8AF3", "5", "Lainnya", "20", 0, y, 115)

	// ===============
	// Section 4
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0046A2")
	dc.Fill()

	imgSeparator3, _ := gg.LoadImage("./assets/sekat_3_360.png")
	dc.DrawImage(imgSeparator3, 0, int(y))
	y += float64(imgSeparator3.Bounds().Dy())
	y += 16

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBold, 16)
	y += drawString(dc, "Paling menyenangkan, selama pakai Raya total keuntungan kamu mencapai", c, y, safeW)
	y += 16

	dc.LoadFontFace(fontBold, 20)
	y += drawStringInBox(dc, "Rp.1.500.000.000", c, y, 8, "#FFC4E8", 8)
	y += 8

	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 12)
	y += drawString(dc, "*Cashback dan Bunga", c, y, safeW)
	y += 4

	// ===============
	// Section 5
	// ===============
	dc.DrawRectangle(0, y, w, h-y)
	dc.SetHexColor("#0D66C5")
	dc.Fill()

	imgSeparator4, _ := gg.LoadImage("./assets/sekat_4_360.png")
	dc.DrawImage(imgSeparator4, 0, int(y))
	y += float64(imgSeparator4.Bounds().Dy())

	dc.SetColor(color.White)
	dc.LoadFontFace(fontBold, 16)
	y += drawString(dc, "Yuk, tingkatkan lagi hubungan ini dengan rutin nabung dan transaksi di Raya biar makin dekat ke depannya!", c, y, 275)
	y += 32

	imgFooter, _ := gg.LoadImage("./assets/footer.png")
	dc.DrawImageAnchored(imgFooter, c, int(y), 0.5, 0)
	y += float64(imgFooter.Bounds().Dy())

	// Export
	dc.SavePNG("output.png")

}

// Draw wrapped string in center and returns its height
func drawString(dc *gg.Context, s string, x float64, y float64, w float64) float64 {
	const lineSpacing = 1.6

	splitLines := dc.WordWrap(s, w)
	_, result := dc.MeasureMultilineString(strings.Join(splitLines, "\n"), lineSpacing)

	dc.DrawStringWrapped(s, x, y, 0.5, 0, w, lineSpacing, gg.AlignCenter)

	return result
}

// Draw string in a rounded rectangle and returns its height
func drawStringInBox(dc *gg.Context, s string, x float64, y float64, r float64, hex string, padding float64) float64 {
	sw, sh := dc.MeasureString(s)

	swp := sw + padding*2
	shp := sh + padding*2

	rX := x - (swp / 2)

	dc.DrawRoundedRectangle(rX, y, swp, shp, r)
	dc.SetHexColor(hex)
	dc.Fill()
	dc.SetColor(color.Black)
	dc.DrawRoundedRectangle(rX, y, swp, shp, r)
	dc.Stroke()

	dc.DrawStringAnchored(s, x, y+(shp/2), 0.5, 0.5)

	return shp
}

// Draw bar with number, name and count. Returns fixed bar height
func drawBarChart(dc *gg.Context, hex string, number string, name string, count string, x float64, y float64, w float64) float64 {
	const barH = 35.0

	dc.DrawRectangle(x, y, w, barH)
	dc.SetHexColor(hex)
	dc.Fill()
	dc.SetColor(color.Black)
	dc.DrawRectangle(x, y, w, barH)
	dc.Stroke()

	cx := x + w - barH/2 - 4
	cy := y + barH/2
	cr := barH/2 - 4

	dc.DrawCircle(cx, cy, cr)
	dc.Stroke()

	dc.LoadFontFace(fontBold, 18)
	dc.DrawStringAnchored(number, cx, cy, 0.5, 0.5)

	nx := x + w + 16
	ny := y
	dc.SetColor(color.White)
	dc.LoadFontFace(fontRegular, 9)
	dc.DrawStringAnchored(name, nx, ny, 0, 1)
	ny += dc.FontHeight()
	ny += 8

	dc.SetHexColor(hex)
	dc.LoadFontFace(fontBlack, 26)
	dc.DrawStringAnchored(count, nx, ny, 0, 1)

	cnw, _ := dc.MeasureString(count)
	nx += cnw + 4
	ny += dc.FontHeight() / 2

	dc.LoadFontFace(fontBold, 13)
	dc.DrawStringAnchored("x", nx, ny, 0, 0.5)

	return barH

}