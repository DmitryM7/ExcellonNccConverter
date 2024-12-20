package main

import "go.uber.org/zap"

/***********************************************
 * Пакет преобразует выгрузку DRL из PCAD 2002 *
 * в формат станка cnc3018                     *
 ***********************************************
 * Решение НЕ универсальное делал под свои     *
 * задачи.                                     *
 ***********************************************
 * G17
 * G20
 * G90
 ***********************************************/
var (
	logger *zap.Logger
	sugar  *zap.SugaredLogger
)

func main() {
	var errLogger error
	logger, errLogger = zap.NewDevelopment()

	if errLogger != nil {
		sugar.Errorln("CAN'T CREATE LOGGER")
	}

	sugar = logger.Sugar()

	parseParams()
}
