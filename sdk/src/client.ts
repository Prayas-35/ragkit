import axios, { type AxiosInstance } from "axios"
import { query, QueryInput, QueryResponse } from "./query"
import { ingest, IngestInput, IngestResponse } from "./ingest"

export class RagClient {
    private client: AxiosInstance

    constructor(config: { apiKey: string; baseUrl?: string }) {
        this.client = axios.create({
            baseURL: config.baseUrl || "https://api.ragkit.dev",
            headers: {
                Authorization: `Bearer ${config.apiKey}`,
                "Content-Type": "application/json",
            },
        })
    }

    async query(input: QueryInput) {
        return query(this.client, input)
    }

    async ingest(input: IngestInput) {
        return ingest(this.client, input)
    }
}