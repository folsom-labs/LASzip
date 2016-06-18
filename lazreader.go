package main

// code related to .laz compressed format

// LazVLRItem describes laz vlr item
type LazVLRItem struct {
	Type    uint16
	Size    uint16
	Version uint16
}

// LazVLR describes laz vlr
type LazVLR struct {
	Compressor      uint16
	Coder           uint16
	VersionMajor    uint8
	VersionMinor    uint8
	VersionRevision uint16
	Options         uint32
	ChunkSize       uint32
	NumPoints       int64
	NumBytes        int64
	NumItems        uint16
	Items           []LazVLRItem
}

// ReadLazVlr reads laz info vlr record
// the data of the LASzip VLR
//     U16  compressor         2 bytes
//     U16  coder              2 bytes
//     U8   version_major      1 byte
//     U8   version_minor      1 byte
//     U16  version_revision   2 bytes
//     U32  options            4 bytes
//     U32  chunk_size         4 bytes
//     I64  num_points         8 bytes
//     I64  num_bytes          8 bytes
//     U16  num_items          2 bytes
//        U16 type                2 bytes * num_items
//        U16 size                2 bytes * num_items
//        U16 version             2 bytes * num_items
// which totals 34+6*num_items
func ReadLazVlr(br *BinaryReader, vlr *VariableLengthRecord) (*LazVLR, error) {
	var res LazVLR

	res.Compressor = br.ReadUint16()
	res.Coder = br.ReadUint16()
	res.VersionMajor = br.ReadUint8()
	res.VersionMinor = br.ReadUint8()
	res.VersionRevision = br.ReadUint16()
	res.Options = br.ReadUint32()
	res.ChunkSize = br.ReadUint32()
	res.NumPoints = br.ReadInt64()
	res.NumBytes = br.ReadInt64()
	res.NumItems = br.ReadUint16()
	for i := 0; i < int(res.NumItems); i++ {
		var item LazVLRItem
		item.Type = br.ReadUint16()
		item.Size = br.ReadUint16()
		item.Version = br.ReadUint16()
		res.Items = append(res.Items, item)
	}

	return &res, br.Error
}
