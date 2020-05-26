package mariadblayer

const (
	// TABLE
	ServerTable						= "device_server_tb d"
	NetworkTable					= "device_network_tb d"
	PartTable						= "device_part_tb d"
	// SELECT
	PageSelectQuery 				= "c1.c_name as manufacture_cd,s1.csub_name as model_cd," +
		 						  	  "c2.c_name as device_type_cd,c3.c_name as ownership_cd," +
									  "c4.c_name as ownership_div_cd,c5.c_name as idc_cd," +
									  "s2.csub_name as rack_cd,cp1.cp_name as user_id,d.*"
	SizeSelectQuery					= "c6.c_name as size_cd"
	CommentSelectQuery				= "cp2.cs_company as comment_register_id, device_comment_tb.*"
	//JOIN
	ManufactureServerJoinQuery 		= "INNER JOIN code_tb AS c1 ON c1.c_type = 'device_server' AND c1.c_idx = d.manufacture_cd"
	ManufactureNetworkJoinQuery 	= "INNER JOIN code_tb AS c1 ON c1.c_type = 'device_network' AND c1.c_idx = d.manufacture_cd"
	ManufacturePartJoinQuery 		= "INNER JOIN code_tb AS c1 ON c1.c_type = 'device_part' AND c1.c_idx = d.manufacture_cd"
	ModelJoinQuery 					= "INNER JOIN code_sub_tb AS s1 ON s1.csub_idx = d.model_cd"
	DeviceTypeServerJoinQuery 		= "INNER JOIN code_tb AS c2 ON c2.c_type = 'device_server' AND c2.c_idx = d.device_type_cd"
	DeviceTypeNetworkJoinQuery 		= "INNER JOIN code_tb AS c2 ON c2.c_type = 'device_network' AND c2.c_idx = d.device_type_cd"
	DeviceTypePartJoinQuery 		= "INNER JOIN code_tb AS c2 ON c2.c_type = 'device_part' AND c2.c_idx = d.device_type_cd"
	OwnershipJoinQuery				= "INNER JOIN code_tb AS c3 ON c3.c_type = 'total' AND c3.c_idx = d.ownership_cd"
	OwnershipDivJoinQuery			= "INNER JOIN code_tb AS c4 ON c4.c_type = 'total' AND c4.c_idx = d.ownership_div_cd"
	IdcJoinQuery					= "INNER JOIN code_tb AS c5 ON c5.c_type = 'total' AND c5.c_idx = d.idc_cd"
	RackJoinQuery					= "INNER JOIN code_sub_tb AS s2 ON s2.csub_idx = d.rack_cd"
	SizeJoinQuery					= "INNER JOIN code_tb AS c6 ON c6.c_type = 'total' AND c6.c_idx = d.size_cd"
	CompanyLeftJoinQuery			= "LEFT OUTER JOIN user_tb u ON u.user_id = d.user_id LEFT OUTER JOIN company_tb cp1 ON cp1.cp_idx = u.cp_idx"
	CompanyAndCommentJoinQuery		= "INNER JOIN company_tb AS cp2 ON cs2.cp_idx = device_comment_tb.comment_register_id"
)
