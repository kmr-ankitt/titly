import type { PageLoad } from './$types';
import { API_BASE_URL } from '$lib/api';

export interface RedirectLoadData {
  shortCode: string;
  targetUrl: string | null;
  backendUrl: string;
  errorStatus: number | null;
  errorMessage: string | null;
}

export const load: PageLoad = async ({ params, fetch }): Promise<RedirectLoadData> => {
  const shortCode = params.url;
  const backendUrl = `${API_BASE_URL}/${shortCode}`;

  try {
    // Attempt fetch to backend short url redirect endpoint
    const res = await fetch(backendUrl, { redirect: 'manual' });

    // Handle standard 302/301 HTTP redirects
    if (res.status === 302 || res.status === 301 || res.type === 'opaqueredirect') {
      const location = res.headers.get('location');
      return {
        shortCode,
        targetUrl: location,
        backendUrl,
        errorStatus: null,
        errorMessage: null
      };
    }

    // Handle 404 Not Found from backend
    if (res.status === 404) {
      const data = await res.json().catch(() => ({}));
      return {
        shortCode,
        targetUrl: null,
        backendUrl,
        errorStatus: 404,
        errorMessage: data.error || 'Short URL not found or has been removed'
      };
    }

    // Handle other HTTP errors
    if (!res.ok) {
      const data = await res.json().catch(() => ({}));
      return {
        shortCode,
        targetUrl: null,
        backendUrl,
        errorStatus: res.status,
        errorMessage: data.error || `Backend server returned status ${res.status}`
      };
    }

    // If fetch followed redirect to destination page automatically
    if (res.url && res.url !== backendUrl) {
      return {
        shortCode,
        targetUrl: res.url,
        backendUrl,
        errorStatus: null,
        errorMessage: null
      };
    }

    // Fallback direct backend redirect URL
    return {
      shortCode,
      targetUrl: backendUrl,
      backendUrl,
      errorStatus: null,
      errorMessage: null
    };
  } catch (e: any) {
    return {
      shortCode,
      targetUrl: null,
      backendUrl,
      errorStatus: 500,
      errorMessage: 'Unable to contact Titly API server'
    };
  }
};
