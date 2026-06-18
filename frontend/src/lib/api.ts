const API_BASE = "http://localhost:8080/api/v1"

export interface ApiResponse<T> {
  success: boolean
  message: string
  data: T
  error?: string
}

export interface Monitor {
  id: string
  url: string
  name?: string
  intervalSeconds: number
  isActive: boolean
  lastCheckedAt: string | null
  createdAt: string
  updatedAt: string
}

export interface LatencyMetric {
  id: string
  monitorId: string
  timestamp: string
  responseTimeMs: number
  statusCode: number
  dnsLookupMs: number | null
  tcpConnectMs: number | null
  ttfbMs: number | null
  isUp: boolean
}

async function request<T>(url: string, options?: RequestInit): Promise<T> {
  const response = await fetch(`${API_BASE}${url}`, {
    ...options,
    headers: {
      "Content-Type": "application/json",
      ...options?.headers,
    },
  })

  const json = await response.json()

  if (!response.ok || !json.success) {
    throw new Error(json.error || json.message || "Request failed")
  }

  return json.data
}

export const api = {
  getMonitors: () => request<Monitor[]>("/monitors"),

  createMonitor: (data: { url: string; intervalSeconds: number; name?: string }) =>
    request<Monitor>("/monitors/create", {
      method: "POST",
      body: JSON.stringify(data),
    }),

  getMonitorMetrics: (id: string) =>
    request<LatencyMetric[]>(`/monitors/${id}/metrics`),

  pingMonitor: (id: string) =>
    request<Monitor>(`/monitors/${id}/ping`, { method: "POST" }),

  updateMonitor: (id: string, data: Partial<Monitor>) =>
    request<Monitor>(`/monitors/${id}`, {
      method: "PATCH",
      body: JSON.stringify(data),
    }),

  deleteMonitor: (id: string) =>
    request<null>(`/monitors/${id}`, { method: "DELETE" }),
}
