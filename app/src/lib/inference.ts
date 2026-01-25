import { BACKEND_URL } from "@/lib/constants";
import { URL } from "url";
import { Garage, Prediction } from "./types";

export async function fetchRecentGarage(): Promise<Garage[]> {
  try {
    const response = await fetch(`${BACKEND_URL}/recent`, {
      next: { revalidate: 0 },
    });
    if (!response.ok) {
      throw new Error("Fetch Latest Garage Status failed.");
    }
    const json = await response.json();
    return json;
  } catch (error) {
    console.error(error);
    return [];
  }
}

export async function fetchGaragePrediction(
  garage_id: number,
  increments: number,
): Promise<Prediction | null> {
  try {
    const url = new URL(`${BACKEND_URL}/predict`);
    url.searchParams.set("garage_id", garage_id.toString());
    url.searchParams.set("increment", increments.toString());
    const response = await fetch(`${url.toString()}`, {
      next: { revalidate: 30 },
    });
    const json = await response.json();
    return { name: json.name, forecast: json.forecast, increments: increments };
  } catch (error) {
    console.error(error);
    return null;
  }
}
