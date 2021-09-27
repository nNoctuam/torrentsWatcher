class Torrent {
  id: number;
  page_url: string;
  file_url: string;
  created_at: string;
  updated_at: string;
  uploaded_at: string | null;
  deleted_at: string | null;
}

class TransmissionTorrent {
  ID: number;
  page_url: string;
  file_url: string;
  created_at: string;
  updated_at: string;
  uploaded_at: string | null;
  deleted_at: string | null;
}

export { Torrent, TransmissionTorrent };
