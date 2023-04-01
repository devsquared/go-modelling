package model

// To help visualize what we are building, take the simple example from the project:
//
//								+------------------------------------+
//								| Simple receiving ASN State Machine |
//								+------------------------------------+
//
//			+----------------------+
//			|     ASN Expected     |				Manual Inventory Removal
//			+----------------------+---------------------------+---------------------+
//			| "Status": "expected" |							 					 |
//			+----------+-----------+											     |
//					   |								         					 |
//				       +- EDI 944 received											 |
//					   |															 |
//                     V															 V
//          +--------------------------+		   						   +-------------------+
//          | ASN received and counted |     Manual Inventory Removal      | Inventory Lost    |
//          +--------------------------+----------------+----------------> +-------------------+
//          |   "Status": "received"   |		   						   | "Status": "lost"  |
//          +--------------------------+		  						   +-------------------+

// Model defines the graphical output for a whole model.
