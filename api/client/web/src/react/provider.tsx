'use client'

import { OpenMeterContext, type OpenMeterConfig } from './context'

export function OpenMeterProvider({
  children,
  ...value
}: OpenMeterConfig & { children?: React.ReactNode }) {
  return (
    <OpenMeterContext.Provider value={value}>
      {children}
    </OpenMeterContext.Provider>
  )
}
