package folder

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"github.com/skip2/go-qrcode"
)

func DeleteFolderLater(folder string, delay time.Duration) {
	go func() {
		<-time.After(delay)
		err := os.RemoveAll(folder)
		if err != nil {
			fmt.Printf("Error al eliminar la carpeta %s después del tiempo: %v\n", folder, err)
		} else {
			fmt.Printf("Carpeta eliminada automáticamente: %s\n", folder)
		}
	}()
}

func GenerateReporteVentasPDF() (string, error) {
	baseFolder := "uuid_storage"
	newUUID := uuid.New().String()
	uuidFolder := filepath.Join(baseFolder, newUUID)

	if err := os.MkdirAll(uuidFolder, os.ModePerm); err != nil {
		return "", fmt.Errorf("error al crear la carpeta %s: %v", uuidFolder, err)
	}

	pdfFileName := filepath.Join(uuidFolder, fmt.Sprintf("%s.pdf", newUUID))

	DeleteFolderLater(uuidFolder, 3*time.Minute)

	return pdfFileName, nil
}


func GenerateQRCode(data string) (string, string, error) {
	baseFolder := "uuid_storage"

	newUUID := uuid.New().String()
	uuidFolder := filepath.Join(baseFolder, newUUID)

	if err := os.MkdirAll(uuidFolder, os.ModePerm); err != nil {
		return "", "", fmt.Errorf("error al crear la carpeta %s: %v", uuidFolder, err)
	}

	qrFileName := filepath.Join(uuidFolder, fmt.Sprintf("%s.png", newUUID))
	err := qrcode.WriteFile(data, qrcode.Medium, 256, qrFileName)
	if err != nil {
		return "", "", fmt.Errorf("error al generar el código QR: %v", err)
	}

	pdfFileName := filepath.Join(uuidFolder, fmt.Sprintf("%s.pdf", newUUID))

	DeleteFolderLater(uuidFolder, 3*time.Minute)

	return qrFileName, pdfFileName, nil
}
