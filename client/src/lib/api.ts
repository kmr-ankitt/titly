import axios from "axios";

export const API_BASE_URL = import.meta.env.VITE_API_URL || "http://localhost:4000";

export interface ShortUrlResponse {
  id?: number;
  long_url?: string;
  short_url: string;
}

export interface ApiError {
  message: string;
  status?: number;
}

export const sendUrl = async (url: string): Promise<ShortUrlResponse> => {
  try {
    const res = await axios.post<ShortUrlResponse>(`${API_BASE_URL}/create-short-url`, {
      long_url: url
    }, {
      timeout: 8000,
      headers: {
        "Content-Type": "application/json"
      }
    });

    if (!res.data || !res.data.short_url) {
      throw new Error("Invalid response received from server");
    }

    return res.data;
  } catch (error) {
    if (axios.isAxiosError(error)) {
      const respData = error.response?.data as { error?: string } | undefined;
      if (respData?.error) {
        throw new Error(respData.error);
      }
      if (error.code === "ECONNABORTED") {
        throw new Error("Server request timed out. Please try again.");
      }
      if (!error.response) {
        throw new Error("Unable to connect to Titly server. Please check if the server is running.");
      }
    }
    throw error instanceof Error ? error : new Error("An unexpected error occurred");
  }
};

export const checkHealth = async (): Promise<boolean> => {
  try {
    const res = await axios.get(`${API_BASE_URL}/`, { timeout: 3000 });
    return res.status === 200;
  } catch {
    return false;
  }
};