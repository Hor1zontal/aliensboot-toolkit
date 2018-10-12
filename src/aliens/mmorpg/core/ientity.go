/*******************************************************************************
 * Copyright (c) 2015, 2017 aliens idea(xiamen) Corporation and others.
 * All rights reserved. 
 * Date:
 *     2018/8/31
 * Contributors:
 *     aliens idea(xiamen) Corporation - initial API and implementation
 *     jialin.he <kylinh@gmail.com>
 *******************************************************************************/
package core



// IEntity declares functions that is defined in Entity
// These functions are mostly component functions
type IEntity interface {
	// Entity Lifetime
	OnInit()       // Called when initializing entity struct, override to initialize entity custom fields
	//OnAttrsReady() // Called when entity attributes are ready.
	//OnCreated()    // Called when entity is just created
	OnDestroy()    // Called when entity is destroying (just before destroy)
	// Migration
	OnMigrateOut() // Called just before entity is migrating out
	OnMigrateIn()  // Called just after entity is migrating in
	// Freeze && Restore
	OnFreeze()   // Called when entity is freezing
	OnRestored() // Called when entity is restored
	// Space Operations
	OnEnterSpace()             // Called when entity leaves space
	OnLeaveSpace(space *Space) // Called when entity enters space
	// Client Notifications
	OnClientConnected()    // Called when Client is connected to entity (become player)
	OnClientDisconnected() // Called when Client disconnected

	DescribeEntityType(desc *EntityDesc) // Define entity attributes in this function
}
