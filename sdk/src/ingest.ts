import { AxiosInstance } from "axios"

export interface IngestMetadata {
    type: "text" | "github" | "url" | "file"
    [key: string]: any
}

export interface IngestInput {
    name: string
    text?: string
    metadata?: IngestMetadata
}

export interface IngestResponse {
    status: string
}

export async function ingest(
    client: AxiosInstance,
    input: IngestInput
): Promise<IngestResponse> {

    if (!input.metadata) {
        throw new Error("metadata is required")
    }

    const res = await client.post("/ingest", {
        name: input.name,
        content: input.text,
        metadata: input.metadata,
    })

    return res.data
}