package main

// RWMutex - тип мьютекса блокирует только операции записи.
// Читать данные могут одновременно несколько горутин.
// Если горутина собирается читать данные, то она вызывает метод RLock().
// Если ей нужно изменить данные, то вызывается метод Lock()
