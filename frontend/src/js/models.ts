class Torrent {
  id: bigint;
  page_url: string;
  file_url: string;
  created_at: string;
  updated_at: string;
  uploaded_at: string | null;
  deleted_at: string | null;
}

module.exports.Torrent = Torrent;
