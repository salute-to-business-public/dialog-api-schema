# Notifications

## Preferences

Notification preferences syncs using [ConfigSync](../proto/config_sync.proto) service.

### Preferences application logic

1. Global parameter which enables notifications is `category.{deviceType}.notification.enabled`
2. User can override global preference for every chat using `account.notification.chat{chatKey}.enabled`

### Preferences parameters

#### `platform`

 - `ios`
 - `android`
 - `sailfish`
 - `web`
 - `generic`
 
#### `deviceType`

 - `mobile`
 - `watch`
 - `tablet`
 - `desktop`
 - `generic`
