# Protocol Documentation
<a name="top"/>

## Table of Contents

- [authentication.proto](#authentication.proto)
    - [AuthSession](#dialog.AuthSession)
    - [ForceReloadContacts](#dialog.ForceReloadContacts)
    - [ForceReloadDialogs](#dialog.ForceReloadDialogs)
    - [ForceReloadField](#dialog.ForceReloadField)
    - [ForceReloadHistory](#dialog.ForceReloadHistory)
    - [RequestCompleteOAuth2](#dialog.RequestCompleteOAuth2)
    - [RequestGetAuthSessions](#dialog.RequestGetAuthSessions)
    - [RequestGetOAuth2Params](#dialog.RequestGetOAuth2Params)
    - [RequestResendCode](#dialog.RequestResendCode)
    - [RequestSendAuthCallObsolete](#dialog.RequestSendAuthCallObsolete)
    - [RequestSendAuthCodeObsolete](#dialog.RequestSendAuthCodeObsolete)
    - [RequestSendCodeByPhoneCall](#dialog.RequestSendCodeByPhoneCall)
    - [RequestSignInObsolete](#dialog.RequestSignInObsolete)
    - [RequestSignOut](#dialog.RequestSignOut)
    - [RequestSignUp](#dialog.RequestSignUp)
    - [RequestSignUpObsolete](#dialog.RequestSignUpObsolete)
    - [RequestStartAnonymousAuth](#dialog.RequestStartAnonymousAuth)
    - [RequestStartEmailAuth](#dialog.RequestStartEmailAuth)
    - [RequestStartPhoneAuth](#dialog.RequestStartPhoneAuth)
    - [RequestStartTokenAuth](#dialog.RequestStartTokenAuth)
    - [RequestStartUsernameAuth](#dialog.RequestStartUsernameAuth)
    - [RequestTerminateAllSessions](#dialog.RequestTerminateAllSessions)
    - [RequestTerminateSession](#dialog.RequestTerminateSession)
    - [RequestValidateCode](#dialog.RequestValidateCode)
    - [RequestValidatePassword](#dialog.RequestValidatePassword)
    - [ResponseAuth](#dialog.ResponseAuth)
    - [ResponseGetAuthSessions](#dialog.ResponseGetAuthSessions)
    - [ResponseGetOAuth2Params](#dialog.ResponseGetOAuth2Params)
    - [ResponseSendAuthCodeObsolete](#dialog.ResponseSendAuthCodeObsolete)
    - [ResponseStartEmailAuth](#dialog.ResponseStartEmailAuth)
    - [ResponseStartPhoneAuth](#dialog.ResponseStartPhoneAuth)
    - [ResponseStartUsernameAuth](#dialog.ResponseStartUsernameAuth)
    - [UpdateForceReloadState](#dialog.UpdateForceReloadState)
  
    - [AuthHolder](#dialog.AuthHolder)
    - [EmailActivationType](#dialog.EmailActivationType)
    - [PhoneActivationType](#dialog.PhoneActivationType)
  
  
    - [Authentication](#dialog.Authentication)
  

- [config_sync.proto](#config_sync.proto)
    - [Parameter](#dialog.Parameter)
    - [RequestEditParameter](#dialog.RequestEditParameter)
    - [RequestGetParameters](#dialog.RequestGetParameters)
    - [ResponseGetParameters](#dialog.ResponseGetParameters)
    - [UpdateParameterChanged](#dialog.UpdateParameterChanged)
  
  
  
    - [ConfigSync](#dialog.ConfigSync)
  

- [contacts.proto](#contacts.proto)
    - [EmailToImport](#dialog.EmailToImport)
    - [PhoneToImport](#dialog.PhoneToImport)
    - [RequestAddContact](#dialog.RequestAddContact)
    - [RequestDeferredImportContacts](#dialog.RequestDeferredImportContacts)
    - [RequestGetContacts](#dialog.RequestGetContacts)
    - [RequestImportContacts](#dialog.RequestImportContacts)
    - [RequestRemoveContact](#dialog.RequestRemoveContact)
    - [RequestSearchContacts](#dialog.RequestSearchContacts)
    - [ResponseDeferredImportContacts](#dialog.ResponseDeferredImportContacts)
    - [ResponseGetContacts](#dialog.ResponseGetContacts)
    - [ResponseImportContacts](#dialog.ResponseImportContacts)
    - [ResponseSearchContacts](#dialog.ResponseSearchContacts)
    - [UpdateContactRegistered](#dialog.UpdateContactRegistered)
    - [UpdateContactsAddTaskSuspended](#dialog.UpdateContactsAddTaskSuspended)
    - [UpdateContactsAdded](#dialog.UpdateContactsAdded)
    - [UpdateContactsRemoved](#dialog.UpdateContactsRemoved)
  
  
  
    - [Contacts](#dialog.Contacts)
  

- [counters.proto](#counters.proto)
    - [AppCounters](#dialog.AppCounters)
    - [UnreadDialog](#dialog.UnreadDialog)
    - [UpdateCountersChanged](#dialog.UpdateCountersChanged)
  
  
  
  

- [crypto.proto](#crypto.proto)
    - [RequestKeyExchange](#dialog.RequestKeyExchange)
    - [ResponseKeyExchange](#dialog.ResponseKeyExchange)
  
  
  
    - [Crypto](#dialog.Crypto)
  

- [definitions.proto](#definitions.proto)
    - [DialogOptions](#dialog.DialogOptions)
    - [UUIDValue](#dialog.UUIDValue)
  
  
    - [File-level Extensions](#definitions.proto-extensions)
  
  

- [device_info.proto](#device_info.proto)
    - [ClientInfo](#dialog.ClientInfo)
    - [RequestNotifyAboutDeviceInfo](#dialog.RequestNotifyAboutDeviceInfo)
  
    - [PlatformType](#dialog.PlatformType)
  
  
    - [DeviceInfo](#dialog.DeviceInfo)
  

- [event_bus.proto](#event_bus.proto)
    - [RequestJoinEventBus](#dialog.RequestJoinEventBus)
    - [RequestKeepAliveEventBus](#dialog.RequestKeepAliveEventBus)
    - [RequestPostToEventBus](#dialog.RequestPostToEventBus)
    - [ResponseJoinEventBus](#dialog.ResponseJoinEventBus)
    - [UpdateEventBusDeviceConnected](#dialog.UpdateEventBusDeviceConnected)
    - [UpdateEventBusDeviceDisconnected](#dialog.UpdateEventBusDeviceDisconnected)
    - [UpdateEventBusDisposed](#dialog.UpdateEventBusDisposed)
    - [UpdateEventBusMessage](#dialog.UpdateEventBusMessage)
  
  
  
    - [EventBus](#dialog.EventBus)
  

- [groups.proto](#groups.proto)
    - [Group](#dialog.Group)
    - [GroupFull](#dialog.GroupFull)
    - [GroupMemberPermission](#dialog.GroupMemberPermission)
    - [Member](#dialog.Member)
    - [RequestCreateGroup](#dialog.RequestCreateGroup)
    - [RequestEditGroupAbout](#dialog.RequestEditGroupAbout)
    - [RequestEditGroupAvatar](#dialog.RequestEditGroupAvatar)
    - [RequestEditGroupTitle](#dialog.RequestEditGroupTitle)
    - [RequestEditGroupTopic](#dialog.RequestEditGroupTopic)
    - [RequestGetGroupInviteUrl](#dialog.RequestGetGroupInviteUrl)
    - [RequestGetGroupInviteUrlBase](#dialog.RequestGetGroupInviteUrlBase)
    - [RequestGetGroupMemberPermissions](#dialog.RequestGetGroupMemberPermissions)
    - [RequestInviteUser](#dialog.RequestInviteUser)
    - [RequestJoinGroup](#dialog.RequestJoinGroup)
    - [RequestJoinGroupByPeer](#dialog.RequestJoinGroupByPeer)
    - [RequestKickUser](#dialog.RequestKickUser)
    - [RequestLeaveGroup](#dialog.RequestLeaveGroup)
    - [RequestLoadFullGroups](#dialog.RequestLoadFullGroups)
    - [RequestLoadMembers](#dialog.RequestLoadMembers)
    - [RequestMakeUserAdmin](#dialog.RequestMakeUserAdmin)
    - [RequestMakeUserAdminObsolete](#dialog.RequestMakeUserAdminObsolete)
    - [RequestRemoveGroupAvatar](#dialog.RequestRemoveGroupAvatar)
    - [RequestRevokeInviteUrl](#dialog.RequestRevokeInviteUrl)
    - [RequestSetGroupShortname](#dialog.RequestSetGroupShortname)
    - [RequestTransferOwnership](#dialog.RequestTransferOwnership)
    - [ResponseCreateGroup](#dialog.ResponseCreateGroup)
    - [ResponseEditGroupAvatar](#dialog.ResponseEditGroupAvatar)
    - [ResponseGetGroupInviteUrlBase](#dialog.ResponseGetGroupInviteUrlBase)
    - [ResponseGetGroupMemberPermissions](#dialog.ResponseGetGroupMemberPermissions)
    - [ResponseInviteUrl](#dialog.ResponseInviteUrl)
    - [ResponseJoinGroup](#dialog.ResponseJoinGroup)
    - [ResponseLoadFullGroups](#dialog.ResponseLoadFullGroups)
    - [ResponseLoadMembers](#dialog.ResponseLoadMembers)
    - [ResponseMakeUserAdminObsolete](#dialog.ResponseMakeUserAdminObsolete)
    - [UpdateGroupAboutChanged](#dialog.UpdateGroupAboutChanged)
    - [UpdateGroupAboutChangedObsolete](#dialog.UpdateGroupAboutChangedObsolete)
    - [UpdateGroupAvatarChanged](#dialog.UpdateGroupAvatarChanged)
    - [UpdateGroupAvatarChangedObsolete](#dialog.UpdateGroupAvatarChangedObsolete)
    - [UpdateGroupCanInviteMembersChanged](#dialog.UpdateGroupCanInviteMembersChanged)
    - [UpdateGroupCanSendMessagesChanged](#dialog.UpdateGroupCanSendMessagesChanged)
    - [UpdateGroupCanViewMembersChanged](#dialog.UpdateGroupCanViewMembersChanged)
    - [UpdateGroupHistoryShared](#dialog.UpdateGroupHistoryShared)
    - [UpdateGroupInviteObsolete](#dialog.UpdateGroupInviteObsolete)
    - [UpdateGroupMemberAdminChanged](#dialog.UpdateGroupMemberAdminChanged)
    - [UpdateGroupMemberChanged](#dialog.UpdateGroupMemberChanged)
    - [UpdateGroupMemberDiff](#dialog.UpdateGroupMemberDiff)
    - [UpdateGroupMemberPermissionsChanged](#dialog.UpdateGroupMemberPermissionsChanged)
    - [UpdateGroupMembersBecameAsync](#dialog.UpdateGroupMembersBecameAsync)
    - [UpdateGroupMembersCountChanged](#dialog.UpdateGroupMembersCountChanged)
    - [UpdateGroupMembersUpdateObsolete](#dialog.UpdateGroupMembersUpdateObsolete)
    - [UpdateGroupMembersUpdated](#dialog.UpdateGroupMembersUpdated)
    - [UpdateGroupOwnerChanged](#dialog.UpdateGroupOwnerChanged)
    - [UpdateGroupShortnameChanged](#dialog.UpdateGroupShortnameChanged)
    - [UpdateGroupTitleChanged](#dialog.UpdateGroupTitleChanged)
    - [UpdateGroupTitleChangedObsolete](#dialog.UpdateGroupTitleChangedObsolete)
    - [UpdateGroupTopicChanged](#dialog.UpdateGroupTopicChanged)
    - [UpdateGroupTopicChangedObsolete](#dialog.UpdateGroupTopicChangedObsolete)
    - [UpdateGroupUserInvitedObsolete](#dialog.UpdateGroupUserInvitedObsolete)
    - [UpdateGroupUserKickObsolete](#dialog.UpdateGroupUserKickObsolete)
    - [UpdateGroupUserLeaveObsolete](#dialog.UpdateGroupUserLeaveObsolete)
  
    - [GroupAdminPermission](#dialog.GroupAdminPermission)
    - [GroupType](#dialog.GroupType)
  
  
    - [Groups](#dialog.Groups)
  

- [integrations.proto](#integrations.proto)
    - [RequestGetIntegrationToken](#dialog.RequestGetIntegrationToken)
    - [RequestRevokeIntegrationToken](#dialog.RequestRevokeIntegrationToken)
    - [ResponseIntegrationToken](#dialog.ResponseIntegrationToken)
  
  
  
    - [Integrations](#dialog.Integrations)
  

- [media_and_files.proto](#media_and_files.proto)
    - [AudioLocation](#dialog.AudioLocation)
    - [Avatar](#dialog.Avatar)
    - [AvatarImage](#dialog.AvatarImage)
    - [Color](#dialog.Color)
    - [FastThumb](#dialog.FastThumb)
    - [FileLocation](#dialog.FileLocation)
    - [FileUrlDescription](#dialog.FileUrlDescription)
    - [HTTPHeader](#dialog.HTTPHeader)
    - [ImageLocation](#dialog.ImageLocation)
    - [PredefinedColor](#dialog.PredefinedColor)
    - [RequestCommitFileUpload](#dialog.RequestCommitFileUpload)
    - [RequestGetFileUploadPartUrl](#dialog.RequestGetFileUploadPartUrl)
    - [RequestGetFileUploadUrl](#dialog.RequestGetFileUploadUrl)
    - [RequestGetFileUrl](#dialog.RequestGetFileUrl)
    - [RequestGetFileUrlBuilder](#dialog.RequestGetFileUrlBuilder)
    - [RequestGetFileUrls](#dialog.RequestGetFileUrls)
    - [ResponseCommitFileUpload](#dialog.ResponseCommitFileUpload)
    - [ResponseGetFileUploadPartUrl](#dialog.ResponseGetFileUploadPartUrl)
    - [ResponseGetFileUploadUrl](#dialog.ResponseGetFileUploadUrl)
    - [ResponseGetFileUrl](#dialog.ResponseGetFileUrl)
    - [ResponseGetFileUrlBuilder](#dialog.ResponseGetFileUrlBuilder)
    - [ResponseGetFileUrls](#dialog.ResponseGetFileUrls)
    - [RgbColor](#dialog.RgbColor)
  
    - [Colors](#dialog.Colors)
  
  
    - [MediaAndFiles](#dialog.MediaAndFiles)
  

- [messaging.proto](#messaging.proto)
    - [AudioMedia](#dialog.AudioMedia)
    - [BinaryMessage](#dialog.BinaryMessage)
    - [DeletedMessage](#dialog.DeletedMessage)
    - [Dialog](#dialog.Dialog)
    - [DialogGroup](#dialog.DialogGroup)
    - [DialogIndex](#dialog.DialogIndex)
    - [DialogShort](#dialog.DialogShort)
    - [DocumentEx](#dialog.DocumentEx)
    - [DocumentExPhoto](#dialog.DocumentExPhoto)
    - [DocumentExVideo](#dialog.DocumentExVideo)
    - [DocumentExVoice](#dialog.DocumentExVoice)
    - [DocumentMessage](#dialog.DocumentMessage)
    - [EmptyMessage](#dialog.EmptyMessage)
    - [HistoryMessage](#dialog.HistoryMessage)
    - [ImageMedia](#dialog.ImageMedia)
    - [InteractiveMedia](#dialog.InteractiveMedia)
    - [InteractiveMediaButton](#dialog.InteractiveMediaButton)
    - [InteractiveMediaConfirm](#dialog.InteractiveMediaConfirm)
    - [InteractiveMediaGroup](#dialog.InteractiveMediaGroup)
    - [InteractiveMediaSelect](#dialog.InteractiveMediaSelect)
    - [InteractiveMediaSelectOption](#dialog.InteractiveMediaSelectOption)
    - [InteractiveMediaTranslation](#dialog.InteractiveMediaTranslation)
    - [InteractiveMediaTranslationGroup](#dialog.InteractiveMediaTranslationGroup)
    - [InteractiveMediaWidget](#dialog.InteractiveMediaWidget)
    - [JsonMessage](#dialog.JsonMessage)
    - [MessageAttributes](#dialog.MessageAttributes)
    - [MessageContent](#dialog.MessageContent)
    - [MessageMedia](#dialog.MessageMedia)
    - [MessageReaction](#dialog.MessageReaction)
    - [ParagraphStyle](#dialog.ParagraphStyle)
    - [PinnedMessages](#dialog.PinnedMessages)
    - [QuotedMessage](#dialog.QuotedMessage)
    - [ReferencedMessages](#dialog.ReferencedMessages)
    - [RequestArchiveChat](#dialog.RequestArchiveChat)
    - [RequestClearChat](#dialog.RequestClearChat)
    - [RequestDeleteChat](#dialog.RequestDeleteChat)
    - [RequestDeleteMessageObsolete](#dialog.RequestDeleteMessageObsolete)
    - [RequestDoInteractiveMediaAction](#dialog.RequestDoInteractiveMediaAction)
    - [RequestFavouriteDialog](#dialog.RequestFavouriteDialog)
    - [RequestFetchDialogIndex](#dialog.RequestFetchDialogIndex)
    - [RequestHideDialog](#dialog.RequestHideDialog)
    - [RequestLoadArchived](#dialog.RequestLoadArchived)
    - [RequestLoadDialogs](#dialog.RequestLoadDialogs)
    - [RequestLoadGroupedDialogs](#dialog.RequestLoadGroupedDialogs)
    - [RequestLoadHistory](#dialog.RequestLoadHistory)
    - [RequestMessageRead](#dialog.RequestMessageRead)
    - [RequestMessageReceived](#dialog.RequestMessageReceived)
    - [RequestMessageRemoveReaction](#dialog.RequestMessageRemoveReaction)
    - [RequestMessageSetReaction](#dialog.RequestMessageSetReaction)
    - [RequestNotifyDialogOpened](#dialog.RequestNotifyDialogOpened)
    - [RequestPinMessage](#dialog.RequestPinMessage)
    - [RequestSendMessage](#dialog.RequestSendMessage)
    - [RequestShowDialog](#dialog.RequestShowDialog)
    - [RequestUnfavouriteDialog](#dialog.RequestUnfavouriteDialog)
    - [RequestUnpinMessage](#dialog.RequestUnpinMessage)
    - [RequestUpdateMessage](#dialog.RequestUpdateMessage)
    - [ResponseDialogsOrder](#dialog.ResponseDialogsOrder)
    - [ResponseFetchDialogIndex](#dialog.ResponseFetchDialogIndex)
    - [ResponseLoadArchived](#dialog.ResponseLoadArchived)
    - [ResponseLoadDialogs](#dialog.ResponseLoadDialogs)
    - [ResponseLoadGroupedDialogs](#dialog.ResponseLoadGroupedDialogs)
    - [ResponseLoadHistory](#dialog.ResponseLoadHistory)
    - [ResponseReactionsResponse](#dialog.ResponseReactionsResponse)
    - [ServiceEx](#dialog.ServiceEx)
    - [ServiceExChangedAbout](#dialog.ServiceExChangedAbout)
    - [ServiceExChangedAvatar](#dialog.ServiceExChangedAvatar)
    - [ServiceExChangedTitle](#dialog.ServiceExChangedTitle)
    - [ServiceExChangedTopic](#dialog.ServiceExChangedTopic)
    - [ServiceExChatArchived](#dialog.ServiceExChatArchived)
    - [ServiceExChatRestored](#dialog.ServiceExChatRestored)
    - [ServiceExContactRegistered](#dialog.ServiceExContactRegistered)
    - [ServiceExGroupCreated](#dialog.ServiceExGroupCreated)
    - [ServiceExPhoneCall](#dialog.ServiceExPhoneCall)
    - [ServiceExPhoneMissed](#dialog.ServiceExPhoneMissed)
    - [ServiceExPhoneRejected](#dialog.ServiceExPhoneRejected)
    - [ServiceExUserInvited](#dialog.ServiceExUserInvited)
    - [ServiceExUserJoined](#dialog.ServiceExUserJoined)
    - [ServiceExUserKicked](#dialog.ServiceExUserKicked)
    - [ServiceExUserLeft](#dialog.ServiceExUserLeft)
    - [ServiceMessage](#dialog.ServiceMessage)
    - [StickerMessage](#dialog.StickerMessage)
    - [TextCommand](#dialog.TextCommand)
    - [TextExMarkdown](#dialog.TextExMarkdown)
    - [TextMessage](#dialog.TextMessage)
    - [TextMessageEx](#dialog.TextMessageEx)
    - [TextModernAttach](#dialog.TextModernAttach)
    - [TextModernField](#dialog.TextModernField)
    - [TextModernMessage](#dialog.TextModernMessage)
    - [UnsupportedMessage](#dialog.UnsupportedMessage)
    - [UpdateChatArchive](#dialog.UpdateChatArchive)
    - [UpdateChatClear](#dialog.UpdateChatClear)
    - [UpdateChatDelete](#dialog.UpdateChatDelete)
    - [UpdateChatGroupsChanged](#dialog.UpdateChatGroupsChanged)
    - [UpdateDialogFavouriteChanged](#dialog.UpdateDialogFavouriteChanged)
    - [UpdateInteractiveMediaEvent](#dialog.UpdateInteractiveMediaEvent)
    - [UpdateMessage](#dialog.UpdateMessage)
    - [UpdateMessageContentChanged](#dialog.UpdateMessageContentChanged)
    - [UpdateMessageDelete](#dialog.UpdateMessageDelete)
    - [UpdateMessageRead](#dialog.UpdateMessageRead)
    - [UpdateMessageReadByMe](#dialog.UpdateMessageReadByMe)
    - [UpdateMessageReceived](#dialog.UpdateMessageReceived)
    - [UpdateMessageSent](#dialog.UpdateMessageSent)
    - [UpdatePinnedMessagesChanged](#dialog.UpdatePinnedMessagesChanged)
    - [UpdateReactionsUpdate](#dialog.UpdateReactionsUpdate)
    - [WebpageMedia](#dialog.WebpageMedia)
  
    - [DialogsFilter](#dialog.DialogsFilter)
    - [InteractiveMediaStyle](#dialog.InteractiveMediaStyle)
    - [ListLoadMode](#dialog.ListLoadMode)
    - [MessageState](#dialog.MessageState)
  
  
    - [Messaging](#dialog.Messaging)
  

- [miscellaneous.proto](#miscellaneous.proto)
    - [Any](#dialog.Any)
    - [CallsConfig](#dialog.CallsConfig)
    - [Config](#dialog.Config)
    - [Discover](#dialog.Discover)
    - [Extension](#dialog.Extension)
    - [InvitesConfig](#dialog.InvitesConfig)
    - [RecursiveMapValue](#dialog.RecursiveMapValue)
    - [RecursiveMapValue.Array](#dialog.RecursiveMapValue.Array)
    - [RecursiveMapValue.Item](#dialog.RecursiveMapValue.Item)
    - [RecursiveMapValue.Value](#dialog.RecursiveMapValue.Value)
    - [ResponseBool](#dialog.ResponseBool)
    - [ResponseSeq](#dialog.ResponseSeq)
    - [ResponseSeqDate](#dialog.ResponseSeqDate)
    - [ResponseSeqDateMid](#dialog.ResponseSeqDateMid)
    - [ResponseVoid](#dialog.ResponseVoid)
    - [ServerMetaInfo](#dialog.ServerMetaInfo)
    - [ServicePeers](#dialog.ServicePeers)
    - [UpdateConfig](#dialog.UpdateConfig)
  
    - [RtcpMuxPolicy](#dialog.RtcpMuxPolicy)
    - [UpdateOptimization](#dialog.UpdateOptimization)
  
  
  

- [obsolete.proto](#obsolete.proto)
    - [ObsoleteGetDifferenceCommand](#dialog.ObsoleteGetDifferenceCommand)
    - [ObsoleteOutPeer](#dialog.ObsoleteOutPeer)
    - [ObsoletePeer](#dialog.ObsoletePeer)
    - [ObsoletePeersList](#dialog.ObsoletePeersList)
    - [ObsoleteSeqUpdateBox](#dialog.ObsoleteSeqUpdateBox)
    - [ObsoleteServiceUpdate](#dialog.ObsoleteServiceUpdate)
    - [ObsoleteWeakUpdateBox](#dialog.ObsoleteWeakUpdateBox)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateTyping](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateTyping)
    - [ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen)
    - [ObsoleteWeakUpdateCommand](#dialog.ObsoleteWeakUpdateCommand)
    - [ObsoleteWeakUpdateCommand.ObsoleteMyOnline](#dialog.ObsoleteWeakUpdateCommand.ObsoleteMyOnline)
    - [ObsoleteWeakUpdateCommand.ObsoleteMyTyping](#dialog.ObsoleteWeakUpdateCommand.ObsoleteMyTyping)
  
    - [ObsoletePeer.ObsoletePeerType](#dialog.ObsoletePeer.ObsoletePeerType)
    - [ObsoleteTypingType](#dialog.ObsoleteTypingType)
  
  
    - [Obsolete](#dialog.Obsolete)
  

- [peers.proto](#peers.proto)
    - [GroupOutPeer](#dialog.GroupOutPeer)
    - [OutPeer](#dialog.OutPeer)
    - [Peer](#dialog.Peer)
    - [UserOutPeer](#dialog.UserOutPeer)
  
    - [PeerType](#dialog.PeerType)
  
  
  

- [privacy.proto](#privacy.proto)
    - [RequestBlockUser](#dialog.RequestBlockUser)
    - [RequestLoadBlockedUsers](#dialog.RequestLoadBlockedUsers)
    - [RequestUnblockUser](#dialog.RequestUnblockUser)
    - [ResponseLoadBlockedUsers](#dialog.ResponseLoadBlockedUsers)
    - [UpdateUserBlocked](#dialog.UpdateUserBlocked)
    - [UpdateUserUnblocked](#dialog.UpdateUserUnblocked)
  
  
  
    - [Privacy](#dialog.Privacy)
  

- [profile.proto](#profile.proto)
    - [RequestChangeUserStatus](#dialog.RequestChangeUserStatus)
    - [RequestCheckNickName](#dialog.RequestCheckNickName)
    - [RequestEditAbout](#dialog.RequestEditAbout)
    - [RequestEditAvatar](#dialog.RequestEditAvatar)
    - [RequestEditCustomProfile](#dialog.RequestEditCustomProfile)
    - [RequestEditMyPreferredLanguages](#dialog.RequestEditMyPreferredLanguages)
    - [RequestEditMyTimeZone](#dialog.RequestEditMyTimeZone)
    - [RequestEditName](#dialog.RequestEditName)
    - [RequestEditNickName](#dialog.RequestEditNickName)
    - [RequestEditSex](#dialog.RequestEditSex)
    - [RequestRemoveAvatar](#dialog.RequestRemoveAvatar)
    - [ResponseEditAvatar](#dialog.ResponseEditAvatar)
  
  
  
    - [Profile](#dialog.Profile)
  

- [push.proto](#push.proto)
    - [RequestRegisterApplePush](#dialog.RequestRegisterApplePush)
    - [RequestRegisterApplePushKit](#dialog.RequestRegisterApplePushKit)
    - [RequestRegisterApplePushToken](#dialog.RequestRegisterApplePushToken)
    - [RequestRegisterGooglePush](#dialog.RequestRegisterGooglePush)
    - [RequestUnregisterApplePush](#dialog.RequestUnregisterApplePush)
    - [RequestUnregisterApplePushKit](#dialog.RequestUnregisterApplePushKit)
    - [RequestUnregisterApplePushToken](#dialog.RequestUnregisterApplePushToken)
    - [RequestUnregisterGooglePush](#dialog.RequestUnregisterGooglePush)
  
  
  
    - [Push](#dialog.Push)
  

- [raw_api.proto](#raw_api.proto)
    - [RequestRawRequest](#dialog.RequestRawRequest)
    - [ResponseRawRequest](#dialog.ResponseRawRequest)
  
  
  
    - [RawAPI](#dialog.RawAPI)
  

- [registration.proto](#registration.proto)
    - [RegisterDeprecatedDeviceRequest](#dialog.RegisterDeprecatedDeviceRequest)
    - [RequestExchangeAuthIdForToken](#dialog.RequestExchangeAuthIdForToken)
    - [RequestRegisterDevice](#dialog.RequestRegisterDevice)
    - [ResponseDeviceRequest](#dialog.ResponseDeviceRequest)
  
  
  
    - [Registration](#dialog.Registration)
  

- [search.proto](#search.proto)
    - [MessageSearchItem](#dialog.MessageSearchItem)
    - [MessageSearchResult](#dialog.MessageSearchResult)
    - [PeerSearchResult](#dialog.PeerSearchResult)
    - [RequestMessageSearch](#dialog.RequestMessageSearch)
    - [RequestMessageSearchMore](#dialog.RequestMessageSearchMore)
    - [RequestPeerSearch](#dialog.RequestPeerSearch)
    - [RequestResolvePeer](#dialog.RequestResolvePeer)
    - [RequestSimpleSearch](#dialog.RequestSimpleSearch)
    - [RequestSimpleSearchMore](#dialog.RequestSimpleSearchMore)
    - [ResponseMessageSearchResponse](#dialog.ResponseMessageSearchResponse)
    - [ResponsePeerSearch](#dialog.ResponsePeerSearch)
    - [ResponseResolvePeer](#dialog.ResponseResolvePeer)
    - [SearchAndCondition](#dialog.SearchAndCondition)
    - [SearchCondition](#dialog.SearchCondition)
    - [SearchOrCondition](#dialog.SearchOrCondition)
    - [SearchPeerCondition](#dialog.SearchPeerCondition)
    - [SearchPeerContentType](#dialog.SearchPeerContentType)
    - [SearchPeerTypeCondition](#dialog.SearchPeerTypeCondition)
    - [SearchPieceText](#dialog.SearchPieceText)
    - [SearchSenderIdConfition](#dialog.SearchSenderIdConfition)
    - [SimpleContactSearchCondition](#dialog.SimpleContactSearchCondition)
    - [SimpleMessageSearchCondition](#dialog.SimpleMessageSearchCondition)
    - [SimplePeerSearchCondition](#dialog.SimplePeerSearchCondition)
    - [SimpleSearchCondition](#dialog.SimpleSearchCondition)
    - [SimpleUserProfileSearchCondition](#dialog.SimpleUserProfileSearchCondition)
    - [criterion](#dialog.criterion)
  
    - [SearchContentType](#dialog.SearchContentType)
    - [SearchPeerType](#dialog.SearchPeerType)
  
  
    - [Search](#dialog.Search)
  

- [sequence_and_updates.proto](#sequence_and_updates.proto)
    - [GroupMembersSubset](#dialog.GroupMembersSubset)
    - [RequestGetDialogsDifference](#dialog.RequestGetDialogsDifference)
    - [RequestGetDifference](#dialog.RequestGetDifference)
    - [RequestGetReferencedEntitites](#dialog.RequestGetReferencedEntitites)
    - [RequestGetState](#dialog.RequestGetState)
    - [RequestSubscribeFromGroupOnline](#dialog.RequestSubscribeFromGroupOnline)
    - [RequestSubscribeFromOnline](#dialog.RequestSubscribeFromOnline)
    - [RequestSubscribeToGroupOnline](#dialog.RequestSubscribeToGroupOnline)
    - [RequestSubscribeToOnline](#dialog.RequestSubscribeToOnline)
    - [ResponseGetDialogsDifference](#dialog.ResponseGetDialogsDifference)
    - [ResponseGetDifference](#dialog.ResponseGetDifference)
    - [ResponseGetReferencedEntitites](#dialog.ResponseGetReferencedEntitites)
    - [SeqUpdateBox](#dialog.SeqUpdateBox)
    - [UpdateCombinedUpdate](#dialog.UpdateCombinedUpdate)
    - [UpdateContainer](#dialog.UpdateContainer)
    - [UpdateEmptyUpdate](#dialog.UpdateEmptyUpdate)
    - [UpdateFatSeqUpdate](#dialog.UpdateFatSeqUpdate)
    - [UpdateRawUpdate](#dialog.UpdateRawUpdate)
    - [UpdateSeqUpdate](#dialog.UpdateSeqUpdate)
    - [UpdateSeqUpdateTooLong](#dialog.UpdateSeqUpdateTooLong)
    - [UpdateWeakFatUpdate](#dialog.UpdateWeakFatUpdate)
    - [UpdateWeakUpdate](#dialog.UpdateWeakUpdate)
  
  
  
    - [SequenceAndUpdates](#dialog.SequenceAndUpdates)
  

- [stickers.proto](#stickers.proto)
    - [RequestAddStickerCollection](#dialog.RequestAddStickerCollection)
    - [RequestAddStickerPackReference](#dialog.RequestAddStickerPackReference)
    - [RequestLoadAcesssibleStickers](#dialog.RequestLoadAcesssibleStickers)
    - [RequestLoadOwnStickers](#dialog.RequestLoadOwnStickers)
    - [RequestLoadStickerCollection](#dialog.RequestLoadStickerCollection)
    - [RequestRemoveStickerCollection](#dialog.RequestRemoveStickerCollection)
    - [RequestRemoveStickerPackReference](#dialog.RequestRemoveStickerPackReference)
    - [ResponseLoadAcesssibleStickers](#dialog.ResponseLoadAcesssibleStickers)
    - [ResponseLoadOwnStickers](#dialog.ResponseLoadOwnStickers)
    - [ResponseLoadStickerCollection](#dialog.ResponseLoadStickerCollection)
    - [ResponseStickersResponse](#dialog.ResponseStickersResponse)
    - [StickerCollection](#dialog.StickerCollection)
    - [StickerDescriptor](#dialog.StickerDescriptor)
    - [UpdateStickerCollectionsChanged](#dialog.UpdateStickerCollectionsChanged)
    - [UpdateStickerPackAdded](#dialog.UpdateStickerPackAdded)
    - [UpdateStickerPackRemoved](#dialog.UpdateStickerPackRemoved)
  
  
  
    - [Stickers](#dialog.Stickers)
  

- [typing_and_online.proto](#typing_and_online.proto)
    - [RequestPauseNotifications](#dialog.RequestPauseNotifications)
    - [RequestRestoreNotifications](#dialog.RequestRestoreNotifications)
    - [RequestSetOnline](#dialog.RequestSetOnline)
    - [RequestStopTyping](#dialog.RequestStopTyping)
    - [RequestTyping](#dialog.RequestTyping)
    - [UpdateGroupOnline](#dialog.UpdateGroupOnline)
    - [UpdatePauseNotifications](#dialog.UpdatePauseNotifications)
    - [UpdateRestoreNotifications](#dialog.UpdateRestoreNotifications)
    - [UpdateTyping](#dialog.UpdateTyping)
    - [UpdateTypingStop](#dialog.UpdateTypingStop)
    - [UpdateUserLastSeen](#dialog.UpdateUserLastSeen)
    - [UpdateUserOffline](#dialog.UpdateUserOffline)
    - [UpdateUserOnline](#dialog.UpdateUserOnline)
  
    - [DeviceType](#dialog.DeviceType)
    - [TypingType](#dialog.TypingType)
  
  
    - [TypingAndOnline](#dialog.TypingAndOnline)
  

- [users.proto](#users.proto)
    - [BotCommand](#dialog.BotCommand)
    - [ContactRecord](#dialog.ContactRecord)
    - [FullUser](#dialog.FullUser)
    - [RequestEditUserLocalName](#dialog.RequestEditUserLocalName)
    - [RequestLoadFullUsers](#dialog.RequestLoadFullUsers)
    - [ResponseLoadFullUsers](#dialog.ResponseLoadFullUsers)
    - [UpdateUserAboutChanged](#dialog.UpdateUserAboutChanged)
    - [UpdateUserAvatarChanged](#dialog.UpdateUserAvatarChanged)
    - [UpdateUserBotCommandsChanged](#dialog.UpdateUserBotCommandsChanged)
    - [UpdateUserContactsChanged](#dialog.UpdateUserContactsChanged)
    - [UpdateUserCustomProfileChanged](#dialog.UpdateUserCustomProfileChanged)
    - [UpdateUserExtChanged](#dialog.UpdateUserExtChanged)
    - [UpdateUserFullExtChanged](#dialog.UpdateUserFullExtChanged)
    - [UpdateUserLocalNameChanged](#dialog.UpdateUserLocalNameChanged)
    - [UpdateUserNameChanged](#dialog.UpdateUserNameChanged)
    - [UpdateUserNickChanged](#dialog.UpdateUserNickChanged)
    - [UpdateUserPreferredLanguagesChanged](#dialog.UpdateUserPreferredLanguagesChanged)
    - [UpdateUserSexChanged](#dialog.UpdateUserSexChanged)
    - [UpdateUserStatusChanged](#dialog.UpdateUserStatusChanged)
    - [UpdateUserTimeZoneChanged](#dialog.UpdateUserTimeZoneChanged)
    - [User](#dialog.User)
    - [UserStatus](#dialog.UserStatus)
  
    - [ContactType](#dialog.ContactType)
    - [Sex](#dialog.Sex)
    - [UserStatusType](#dialog.UserStatusType)
  
  
    - [Users](#dialog.Users)
  

- [web_rtc.proto](#web_rtc.proto)
    - [ActiveCall](#dialog.ActiveCall)
    - [AdvertiseMaster](#dialog.AdvertiseMaster)
    - [AdvertisePeer](#dialog.AdvertisePeer)
    - [AdvertiseSelf](#dialog.AdvertiseSelf)
    - [Answer](#dialog.Answer)
    - [CallLogEntry](#dialog.CallLogEntry)
    - [CallMember](#dialog.CallMember)
    - [CallMemberStateHolder](#dialog.CallMemberStateHolder)
    - [CallNameChanged](#dialog.CallNameChanged)
    - [CallStats](#dialog.CallStats)
    - [Candidate](#dialog.Candidate)
    - [CloseSession](#dialog.CloseSession)
    - [DTMF](#dialog.DTMF)
    - [EnableConnection](#dialog.EnableConnection)
    - [GotICECandidate](#dialog.GotICECandidate)
    - [ICECandidate](#dialog.ICECandidate)
    - [ICEServer](#dialog.ICEServer)
    - [NeedDisconnect](#dialog.NeedDisconnect)
    - [NeedOffer](#dialog.NeedOffer)
    - [NegotinationSuccessful](#dialog.NegotinationSuccessful)
    - [Offer](#dialog.Offer)
    - [OnRenegotiationNeeded](#dialog.OnRenegotiationNeeded)
    - [PeerSettings](#dialog.PeerSettings)
    - [RemovedICECandidates](#dialog.RemovedICECandidates)
    - [RequestChangeCallDisplayName](#dialog.RequestChangeCallDisplayName)
    - [RequestDeleteCall](#dialog.RequestDeleteCall)
    - [RequestDoCall](#dialog.RequestDoCall)
    - [RequestGetCallInfo](#dialog.RequestGetCallInfo)
    - [RequestJoinCall](#dialog.RequestJoinCall)
    - [RequestLoadCalls](#dialog.RequestLoadCalls)
    - [RequestRejectCall](#dialog.RequestRejectCall)
    - [ResponseDoCall](#dialog.ResponseDoCall)
    - [ResponseGetCallInfo](#dialog.ResponseGetCallInfo)
    - [ResponseLoadCalls](#dialog.ResponseLoadCalls)
    - [UpdateCallDisposed](#dialog.UpdateCallDisposed)
    - [UpdateCallHandled](#dialog.UpdateCallHandled)
    - [UpdateIncomingCall](#dialog.UpdateIncomingCall)
    - [UpdateIncomingCallDeprecated](#dialog.UpdateIncomingCallDeprecated)
    - [WebRTCSignaling](#dialog.WebRTCSignaling)
  
    - [CallDisposedReason](#dialog.CallDisposedReason)
    - [CallMemberState](#dialog.CallMemberState)
    - [CallStatsType](#dialog.CallStatsType)
    - [DTMFCode](#dialog.DTMFCode)
    - [RejectCallReason](#dialog.RejectCallReason)
  
  
    - [WebRTC](#dialog.WebRTC)
  

- [Scalar Value Types](#scalar-value-types)



<a name="authentication.proto"/>
<p align="right"><a href="#top">Top</a></p>

## authentication.proto



<a name="dialog.AuthSession"/>

### AuthSession
Authentication session


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Unuque ID of session |
| auth_holder | [AuthHolder](#dialog.AuthHolder) |  | holder of session. 1 - this device, 2 - other. |
| app_id | [int32](#int32) |  | Application Id that you set during authorization |
| app_title | [string](#string) |  | Deprecated |
| device_title | [string](#string) |  | Deprecated |
| auth_time | [int32](#int32) |  | Time of session creating |
| auth_location | [string](#string) |  | Two-letter country code of session create |
| latitude | [google.protobuf.DoubleValue](#google.protobuf.DoubleValue) |  |  |
| longitude | [google.protobuf.DoubleValue](#google.protobuf.DoubleValue) |  |  |






<a name="dialog.ForceReloadContacts"/>

### ForceReloadContacts
Tells the client to clear contacts and load them again






<a name="dialog.ForceReloadDialogs"/>

### ForceReloadDialogs
Tells the client to clear dialogs and load them again






<a name="dialog.ForceReloadField"/>

### ForceReloadField
Notification to force client to reload some entities from server
Just for old clients. Should be ignore.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| forceReloadDialogs | [ForceReloadDialogs](#dialog.ForceReloadDialogs) |  |  |
| forceReloadContacts | [ForceReloadContacts](#dialog.ForceReloadContacts) |  |  |
| forceReloadHistory | [ForceReloadHistory](#dialog.ForceReloadHistory) |  |  |






<a name="dialog.ForceReloadHistory"/>

### ForceReloadHistory
Tells the client to clear the specified conversation and load it again
peer the peer whose history should be reloaded


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |






<a name="dialog.RequestCompleteOAuth2"/>

### RequestCompleteOAuth2
Complete OAuth2 Authentication - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  |  |
| code | [string](#string) |  |  |






<a name="dialog.RequestGetAuthSessions"/>

### RequestGetAuthSessions
Getting of all active user&#39;s authentication sessions






<a name="dialog.RequestGetOAuth2Params"/>

### RequestGetOAuth2Params
Loading OAuth2 Parameters - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  |  |
| redirect_url | [string](#string) |  |  |






<a name="dialog.RequestResendCode"/>

### RequestResendCode
Performs code resend


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of the authorization transaction * |






<a name="dialog.RequestSendAuthCallObsolete"/>

### RequestSendAuthCallObsolete
Requesting Phone activation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  |  |
| sms_hash | [string](#string) |  |  |
| app_id | [int32](#int32) |  |  |
| api_key | [string](#string) |  |  |






<a name="dialog.RequestSendAuthCodeObsolete"/>

### RequestSendAuthCodeObsolete
Sending SMS with activation code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  |  |
| app_id | [int32](#int32) |  |  |
| api_key | [string](#string) |  |  |






<a name="dialog.RequestSendCodeByPhoneCall"/>

### RequestSendCodeByPhoneCall
Dial phone and dictate auth code


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash from ResponseStartPhoneAuth |






<a name="dialog.RequestSignInObsolete"/>

### RequestSignInObsolete
Performing user signin - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  |  |
| sms_hash | [string](#string) |  |  |
| sms_code | [string](#string) |  |  |
| device_hash | [bytes](#bytes) |  |  |
| device_title | [string](#string) |  |  |
| app_id | [int32](#int32) |  |  |
| app_key | [string](#string) |  |  |






<a name="dialog.RequestSignOut"/>

### RequestSignOut
SignOut current session






<a name="dialog.RequestSignUp"/>

### RequestSignUp
Perform user SignUp


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of the authorization transaction * |
| name | [string](#string) |  | Your name |
| sex | [Sex](#dialog.Sex) |  | Use it in case of anonymous authorization (deprecated) |
| password | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestSignUpObsolete"/>

### RequestSignUpObsolete
Performing user signup. If user perform signup on already registered user it just override previous
profile information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  |  |
| sms_hash | [string](#string) |  |  |
| sms_code | [string](#string) |  |  |
| name | [string](#string) |  |  |
| device_hash | [bytes](#bytes) |  |  |
| device_title | [string](#string) |  |  |
| app_id | [int32](#int32) |  |  |
| app_key | [string](#string) |  |  |
| is_silent | [bool](#bool) |  |  |






<a name="dialog.RequestStartAnonymousAuth"/>

### RequestStartAnonymousAuth
Starting Anonymous login - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| app_id | [int32](#int32) |  |  |
| api_key | [string](#string) |  |  |
| device_hash | [bytes](#bytes) |  |  |
| device_title | [string](#string) |  |  |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| preferred_languages | [string](#string) | repeated |  |






<a name="dialog.RequestStartEmailAuth"/>

### RequestStartEmailAuth
Start EMail Activation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  |  |
| app_id | [int32](#int32) |  | Application id, choose it by yourself and hold during authorization process * |
| api_key | [string](#string) |  | Deprecated field - keep it empty * |
| device_hash | [bytes](#bytes) |  | Deprecated field - keep it empty * |
| device_title | [string](#string) |  | Some title, choose it by yourself and hold during authorization process * |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Your timezone * |
| preferred_languages | [string](#string) | repeated | First language from this array will be used for some notifications from server * |






<a name="dialog.RequestStartPhoneAuth"/>

### RequestStartPhoneAuth
Start Phone Activation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  |  |
| app_id | [int32](#int32) |  | Application id, choose it by yourself and hold during authorization process * |
| api_key | [string](#string) |  | Deprecated field - keep it empty * |
| device_hash | [bytes](#bytes) |  | Deprecated field - keep it empty * |
| device_title | [string](#string) |  | Some title, choose it by yourself and hold during authorization process * |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Your timezone * |
| preferred_languages | [string](#string) | repeated | First language from this array will be used for some notifications from server * |






<a name="dialog.RequestStartTokenAuth"/>

### RequestStartTokenAuth
Starting token-based login - to authorize bot


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  | Token received from BotFather * |
| app_id | [int32](#int32) |  | Application id, choose it by yourself and hold during authorization process * |
| api_key | [string](#string) |  | Deprecated field - keep it empty * |
| device_hash | [bytes](#bytes) |  | Deprecated field - keep it empty * |
| device_title | [string](#string) |  | Some title, choose it by yourself and hold during authorization process * |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Your timezone * |
| preferred_languages | [string](#string) | repeated | First language from this array will be used for some notifications from server * |






<a name="dialog.RequestStartUsernameAuth"/>

### RequestStartUsernameAuth
Starting Login Authentication


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  |  |
| app_id | [int32](#int32) |  | Application id, choose it by yourself and hold during authorization process * |
| api_key | [string](#string) |  | Deprecated field - keep it empty * |
| device_hash | [bytes](#bytes) |  | Deprecated field - keep it empty * |
| device_title | [string](#string) |  | Some title, choose it by yourself and hold during authorization process * |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| preferred_languages | [string](#string) | repeated | First language from this array will be used for some notifications from server * |






<a name="dialog.RequestTerminateAllSessions"/>

### RequestTerminateAllSessions
SignOut on all exept current sessions






<a name="dialog.RequestTerminateSession"/>

### RequestTerminateSession
SignOut on specified user&#39;s session


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Session id |






<a name="dialog.RequestValidateCode"/>

### RequestValidateCode
Performing user sign in.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of the authorization transaction * |
| code | [string](#string) |  | Authorization code |






<a name="dialog.RequestValidatePassword"/>

### RequestValidatePassword
Validation of account password


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of the authorization transaction * |
| password | [string](#string) |  |  |






<a name="dialog.ResponseAuth"/>

### ResponseAuth
Authentication result


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user | [User](#dialog.User) |  | Registered/authorized user |
| config | [Config](#dialog.Config) |  | Config for that user |
| config_hash | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  | Hash of config to later usage |






<a name="dialog.ResponseGetAuthSessions"/>

### ResponseGetAuthSessions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_auths | [AuthSession](#dialog.AuthSession) | repeated | Active user&#39;s sessions |






<a name="dialog.ResponseGetOAuth2Params"/>

### ResponseGetOAuth2Params
Deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth_url | [string](#string) |  |  |






<a name="dialog.ResponseSendAuthCodeObsolete"/>

### ResponseSendAuthCodeObsolete



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sms_hash | [string](#string) |  |  |
| is_registered | [bool](#bool) |  |  |






<a name="dialog.ResponseStartEmailAuth"/>

### ResponseStartEmailAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of authorization transaction * |
| is_registered | [bool](#bool) |  | Deprecated |
| activation_type | [EmailActivationType](#dialog.EmailActivationType) |  | Code or password - call ValidateCode if code * |






<a name="dialog.ResponseStartPhoneAuth"/>

### ResponseStartPhoneAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of authorization transaction * |
| is_registered | [bool](#bool) |  | Deprecated |
| activation_type | [PhoneActivationType](#dialog.PhoneActivationType) |  | Code or password - call ValidateCode if code * |






<a name="dialog.ResponseStartUsernameAuth"/>

### ResponseStartUsernameAuth



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| transaction_hash | [string](#string) |  | Hash of authorization transaction * |
| is_registered | [bool](#bool) |  | Deprecated |






<a name="dialog.UpdateForceReloadState"/>

### UpdateForceReloadState
This update is sent by the server to force a client to reload its data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [ForceReloadField](#dialog.ForceReloadField) | repeated |  |





 


<a name="dialog.AuthHolder"/>

### AuthHolder
Holder of session

| Name | Number | Description |
| ---- | ------ | ----------- |
| AUTHHOLDER_UNKNOWN | 0 |  |
| AUTHHOLDER_THISDEVICE | 1 |  |
| AUTHHOLDER_OTHERDEVICE | 2 |  |



<a name="dialog.EmailActivationType"/>

### EmailActivationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EMAILACTIVATIONTYPE_UNKNOWN | 0 |  |
| EMAILACTIVATIONTYPE_CODE | 1 |  |
| EMAILACTIVATIONTYPE_OAUTH2 | 2 |  |
| EMAILACTIVATIONTYPE_PASSWORD | 3 |  |



<a name="dialog.PhoneActivationType"/>

### PhoneActivationType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PHONEACTIVATIONTYPE_UNKNOWN | 0 |  |
| PHONEACTIVATIONTYPE_CODE | 1 |  |
| PHONEACTIVATIONTYPE_PASSWORD | 2 |  |


 

 


<a name="dialog.Authentication"/>

### Authentication


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| StartPhoneAuth | [RequestStartPhoneAuth](#dialog.RequestStartPhoneAuth) | [ResponseStartPhoneAuth](#dialog.RequestStartPhoneAuth) | Start authorization by phone |
| SendCodeByPhoneCall | [RequestSendCodeByPhoneCall](#dialog.RequestSendCodeByPhoneCall) | [ResponseVoid](#dialog.RequestSendCodeByPhoneCall) | Resend code by transaction hash |
| StartEmailAuth | [RequestStartEmailAuth](#dialog.RequestStartEmailAuth) | [ResponseStartEmailAuth](#dialog.RequestStartEmailAuth) | Start email authorization process |
| StartAnonymousAuth | [RequestStartAnonymousAuth](#dialog.RequestStartAnonymousAuth) | [ResponseAuth](#dialog.RequestStartAnonymousAuth) | Deprecated |
| StartTokenAuth | [RequestStartTokenAuth](#dialog.RequestStartTokenAuth) | [ResponseAuth](#dialog.RequestStartTokenAuth) | Start token auth authorization (actual for bots) |
| StartUsernameAuth | [RequestStartUsernameAuth](#dialog.RequestStartUsernameAuth) | [ResponseStartUsernameAuth](#dialog.RequestStartUsernameAuth) | Start login/password authorization process |
| ValidateCode | [RequestValidateCode](#dialog.RequestValidateCode) | [ResponseAuth](#dialog.RequestValidateCode) | Validate code received by phone or email Returns error if user does not exist |
| ResendCode | [RequestResendCode](#dialog.RequestResendCode) | [ResponseVoid](#dialog.RequestResendCode) | Resend code if you don&#39;t receive it with first attempt |
| ValidatePassword | [RequestValidatePassword](#dialog.RequestValidatePassword) | [ResponseAuth](#dialog.RequestValidatePassword) | Validate your passwword |
| GetOAuth2Params | [RequestGetOAuth2Params](#dialog.RequestGetOAuth2Params) | [ResponseGetOAuth2Params](#dialog.RequestGetOAuth2Params) | Deprecated |
| CompleteOAuth2 | [RequestCompleteOAuth2](#dialog.RequestCompleteOAuth2) | [ResponseAuth](#dialog.RequestCompleteOAuth2) | Deprecated |
| SignUp | [RequestSignUp](#dialog.RequestSignUp) | [ResponseAuth](#dialog.RequestSignUp) | Sign up existed user |
| GetAuthSessions | [RequestGetAuthSessions](#dialog.RequestGetAuthSessions) | [ResponseGetAuthSessions](#dialog.RequestGetAuthSessions) | Returns all authorized user&#39;s sessions |
| TerminateSession | [RequestTerminateSession](#dialog.RequestTerminateSession) | [ResponseVoid](#dialog.RequestTerminateSession) | Deprecated. Does not produce any effect. |
| TerminateAllSessions | [RequestTerminateAllSessions](#dialog.RequestTerminateAllSessions) | [ResponseVoid](#dialog.RequestTerminateAllSessions) | Log out user |
| SignOut | [RequestSignOut](#dialog.RequestSignOut) | [ResponseVoid](#dialog.RequestSignOut) | Log out current session |
| SignInObsolete | [RequestSignInObsolete](#dialog.RequestSignInObsolete) | [ResponseAuth](#dialog.RequestSignInObsolete) | Deprecated |
| SignUpObsolete | [RequestSignUpObsolete](#dialog.RequestSignUpObsolete) | [ResponseAuth](#dialog.RequestSignUpObsolete) | Deprecated |
| SendAuthCodeObsolete | [RequestSendAuthCodeObsolete](#dialog.RequestSendAuthCodeObsolete) | [ResponseSendAuthCodeObsolete](#dialog.RequestSendAuthCodeObsolete) | Deprecated |
| SendAuthCallObsolete | [RequestSendAuthCallObsolete](#dialog.RequestSendAuthCallObsolete) | [ResponseVoid](#dialog.RequestSendAuthCallObsolete) | Deprecated |

 



<a name="config_sync.proto"/>
<p align="right"><a href="#top">Top</a></p>

## config_sync.proto



<a name="dialog.Parameter"/>

### Parameter
Syncing Parameter


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | Key of parameter |
| value | [string](#string) |  | Value of parameter |






<a name="dialog.RequestEditParameter"/>

### RequestEditParameter
Change parameter value


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestGetParameters"/>

### RequestGetParameters
Getting Parameters






<a name="dialog.ResponseGetParameters"/>

### ResponseGetParameters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| parameters | [Parameter](#dialog.Parameter) | repeated |  |






<a name="dialog.UpdateParameterChanged"/>

### UpdateParameterChanged
Update about parameter change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |





 

 

 


<a name="dialog.ConfigSync"/>

### ConfigSync


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetParameters | [RequestGetParameters](#dialog.RequestGetParameters) | [ResponseGetParameters](#dialog.RequestGetParameters) |  |
| EditParameter | [RequestEditParameter](#dialog.RequestEditParameter) | [ResponseSeq](#dialog.RequestEditParameter) |  |

 



<a name="contacts.proto"/>
<p align="right"><a href="#top">Top</a></p>

## contacts.proto



<a name="dialog.EmailToImport"/>

### EmailToImport
Email for import


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | email for importing |
| name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | optional name for contact |






<a name="dialog.PhoneToImport"/>

### PhoneToImport
Phone for import


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phone_number | [int64](#int64) |  | phone number for import in international format |
| name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | optional name for contact |






<a name="dialog.RequestAddContact"/>

### RequestAddContact
Adding contact to contact list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| access_hash | [int64](#int64) |  |  |






<a name="dialog.RequestDeferredImportContacts"/>

### RequestDeferredImportContacts
Importing phones and emails for building contact list
Import evaluated lazily, response does not contain any info
Maximum amount of items for import per method call equals to 100.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phones | [PhoneToImport](#dialog.PhoneToImport) | repeated |  |
| emails | [EmailToImport](#dialog.EmailToImport) | repeated |  |






<a name="dialog.RequestGetContacts"/>

### RequestGetContacts
Getting current contact list
SHA256 hash of list of a comma-separated list of contact UIDs in ascending
order may be passed in contactsHash parameter.
If the contact list was not changed, isNotChanged will be true.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contacts_hash | [string](#string) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestImportContacts"/>

### RequestImportContacts
Importing phones and emails for building contact list
Maximum amount of items for import per method call equals to 100.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| phones | [PhoneToImport](#dialog.PhoneToImport) | repeated |  |
| emails | [EmailToImport](#dialog.EmailToImport) | repeated |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestRemoveContact"/>

### RequestRemoveContact
Removing contact from contact list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| access_hash | [int64](#int64) |  |  |






<a name="dialog.RequestSearchContacts"/>

### RequestSearchContacts
Searching contacts by user&#39;s query


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| request | [string](#string) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.ResponseDeferredImportContacts"/>

### ResponseDeferredImportContacts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [string](#string) |  |  |






<a name="dialog.ResponseGetContacts"/>

### ResponseGetContacts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#dialog.User) | repeated |  |
| is_not_changed | [bool](#bool) |  |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.ResponseImportContacts"/>

### ResponseImportContacts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#dialog.User) | repeated | Registered contacts |
| seq | [int32](#int32) |  | Deprecated |
| state | [bytes](#bytes) |  | Server state related to current client, used by server only |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.ResponseSearchContacts"/>

### ResponseSearchContacts



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#dialog.User) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.UpdateContactRegistered"/>

### UpdateContactRegistered
Update about contact registration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  | User id of a registered contact |
| is_silent | [bool](#bool) |  | deprecated |
| date | [int64](#int64) |  | deprecated |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |






<a name="dialog.UpdateContactsAddTaskSuspended"/>

### UpdateContactsAddTaskSuspended
Update about suspending task - normally it should be ignored


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| task_id | [string](#string) |  |  |






<a name="dialog.UpdateContactsAdded"/>

### UpdateContactsAdded
Update about contacts added


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [int32](#int32) | repeated | User ids of the registered contacts |
| task_id | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Id of the task that finished |






<a name="dialog.UpdateContactsRemoved"/>

### UpdateContactsRemoved
Update about contacts removed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uids | [int32](#int32) | repeated |  |





 

 

 


<a name="dialog.Contacts"/>

### Contacts


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ImportContacts | [RequestImportContacts](#dialog.RequestImportContacts) | [ResponseImportContacts](#dialog.RequestImportContacts) | Import contacts and wait while query is not finished |
| DeferredImportContacts | [RequestDeferredImportContacts](#dialog.RequestDeferredImportContacts) | [ResponseDeferredImportContacts](#dialog.RequestDeferredImportContacts) | Same as above, but without waiting response |
| GetContacts | [RequestGetContacts](#dialog.RequestGetContacts) | [ResponseGetContacts](#dialog.RequestGetContacts) |  |
| RemoveContact | [RequestRemoveContact](#dialog.RequestRemoveContact) | [ResponseSeq](#dialog.RequestRemoveContact) |  |
| AddContact | [RequestAddContact](#dialog.RequestAddContact) | [ResponseSeq](#dialog.RequestAddContact) |  |
| SearchContacts | [RequestSearchContacts](#dialog.RequestSearchContacts) | [ResponseSearchContacts](#dialog.RequestSearchContacts) | Search contacts by query string |

 



<a name="counters.proto"/>
<p align="right"><a href="#top">Top</a></p>

## counters.proto



<a name="dialog.AppCounters"/>

### AppCounters



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| global_counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | Global unread counter |
| global_dialogs_counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | Global count of dialogs with positive counter values |
| unread_dialogs | [UnreadDialog](#dialog.UnreadDialog) | repeated | map of chat peer id to unread to counters |






<a name="dialog.UnreadDialog"/>

### UnreadDialog
Unread dialogs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  | dialog peer |
| counter | [int32](#int32) |  | unread messages count |






<a name="dialog.UpdateCountersChanged"/>

### UpdateCountersChanged
deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| counters | [AppCounters](#dialog.AppCounters) |  |  |
| ts | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |





 

 

 

 



<a name="crypto.proto"/>
<p align="right"><a href="#top">Top</a></p>

## crypto.proto



<a name="dialog.RequestKeyExchange"/>

### RequestKeyExchange
Exchange public keys


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_key | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ResponseKeyExchange"/>

### ResponseKeyExchange



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server_key | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |





 

 

 


<a name="dialog.Crypto"/>

### Crypto


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| KeyExchange | [RequestKeyExchange](#dialog.RequestKeyExchange) | [ResponseKeyExchange](#dialog.RequestKeyExchange) |  |

 



<a name="definitions.proto"/>
<p align="right"><a href="#top">Top</a></p>

## definitions.proto



<a name="dialog.DialogOptions"/>

### DialogOptions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| log | [string](#string) |  |  |






<a name="dialog.UUIDValue"/>

### UUIDValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bytes](#bytes) |  |  |





 

 


<a name="definitions.proto-extensions"/>

### File-level Extensions
| Extension | Type | Base | Number | Description |
| --------- | ---- | ---- | ------ | ----------- |
| dlg | DialogOptions | .google.protobuf.FieldOptions | 100001 |  |

 

 



<a name="device_info.proto"/>
<p align="right"><a href="#top">Top</a></p>

## device_info.proto



<a name="dialog.ClientInfo"/>

### ClientInfo
Generic client info, containing information about platform type, version, sdk etc


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| platform | [PlatformType](#dialog.PlatformType) |  | The platform enum. Can be either Android, Web or iOS |
| device_name | [string](#string) |  | For android: vendor and model; For iOS: model; For Web: platform and user agent |
| app_name | [string](#string) |  | Optional application name |
| app_version | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Application version |
| sdk_version | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Optional SDK version |
| preferred_languages | [string](#string) | repeated | Optional ISO-639 language code and ISO-3166 country code: ru-RU |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Optional TimeZone id |






<a name="dialog.RequestNotifyAboutDeviceInfo"/>

### RequestNotifyAboutDeviceInfo
Notifying about device information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| preferred_languages | [string](#string) | repeated | First language from this array will be used for some notifications from server * |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Your timezone * |





 


<a name="dialog.PlatformType"/>

### PlatformType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PLATFORMTYPE_UNKNOWN | 0 |  |
| PLATFORMTYPE_ANDROID | 1 |  |
| PLATFORMTYPE_IOS | 2 |  |
| PLATFORMTYPE_WEB | 3 |  |
| PLATFORMTYPE_CLC | 4 |  |
| PLATFORMTYPE_TESTS | 42 |  |


 

 


<a name="dialog.DeviceInfo"/>

### DeviceInfo


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| NotifyAboutDeviceInfo | [RequestNotifyAboutDeviceInfo](#dialog.RequestNotifyAboutDeviceInfo) | [ResponseVoid](#dialog.RequestNotifyAboutDeviceInfo) | Set info about current device |

 



<a name="event_bus.proto"/>
<p align="right"><a href="#top">Top</a></p>

## event_bus.proto



<a name="dialog.RequestJoinEventBus"/>

### RequestJoinEventBus
Joining Event Bus


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| timeout | [int64](#int64) |  | Join TTL |






<a name="dialog.RequestKeepAliveEventBus"/>

### RequestKeepAliveEventBus
Keep Alive Event Bus


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| timeout | [int64](#int64) |  | Keep TTL |






<a name="dialog.RequestPostToEventBus"/>

### RequestPostToEventBus
Event Bus Destination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| destinations | [int64](#int64) | repeated | Destination device ids |
| message | [bytes](#bytes) |  | Message to send |






<a name="dialog.ResponseJoinEventBus"/>

### ResponseJoinEventBus



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device_id | [int64](#int64) |  | Generated device id for this event bus |






<a name="dialog.UpdateEventBusDeviceConnected"/>

### UpdateEventBusDeviceConnected
Update about pubsub device connected


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| user_id | [int32](#int32) |  | Connected user id |
| device_id | [int64](#int64) |  | Connected device id |
| peer | [Peer](#dialog.Peer) |  | Connected peer |






<a name="dialog.UpdateEventBusDeviceDisconnected"/>

### UpdateEventBusDeviceDisconnected
Update about device disconnected


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| user_id | [int32](#int32) |  |  |
| device_id | [int64](#int64) |  |  |
| peer | [Peer](#dialog.Peer) |  |  |






<a name="dialog.UpdateEventBusDisposed"/>

### UpdateEventBusDisposed
Event Bus dispose


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |






<a name="dialog.UpdateEventBusMessage"/>

### UpdateEventBusMessage
Event Bus Message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | Bus id |
| sender_id | [int32](#int32) |  |  |
| sender_device_id | [int64](#int64) |  |  |
| message | [bytes](#bytes) |  |  |





 

 

 


<a name="dialog.EventBus"/>

### EventBus


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| JoinEventBus | [RequestJoinEventBus](#dialog.RequestJoinEventBus) | [ResponseJoinEventBus](#dialog.RequestJoinEventBus) |  |
| KeepAliveEventBus | [RequestKeepAliveEventBus](#dialog.RequestKeepAliveEventBus) | [ResponseVoid](#dialog.RequestKeepAliveEventBus) |  |
| PostToEventBus | [RequestPostToEventBus](#dialog.RequestPostToEventBus) | [ResponseVoid](#dialog.RequestPostToEventBus) |  |

 



<a name="groups.proto"/>
<p align="right"><a href="#top">Top</a></p>

## groups.proto



<a name="dialog.Group"/>

### Group
Group information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | group id |
| access_hash | [int64](#int64) |  | Access hash of group |
| title | [string](#string) |  | Title of group |
| avatar | [Avatar](#dialog.Avatar) |  | Avatar of group |
| members_amount | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | Number of members |
| is_member | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is current user a member of a group. Default is true. |
| is_hidden | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is group hidden (not showing it in recent list). Default is false. |
| group_type | [GroupType](#dialog.GroupType) |  | Group Type. Used only for displaying information. Default is GROUP. |
| can_send_message | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Can user send messages. Default is equals isMember for Group and false for channels. |
| is_admin | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Deprecated |
| creator_uid | [int32](#int32) |  | Group creator |
| members | [Member](#dialog.Member) | repeated | Members of group |
| create_date | [int64](#int64) |  | Date of creation |
| theme | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Theme of group |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | About of group |
| shortname | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Group short name |






<a name="dialog.GroupFull"/>

### GroupFull
Goup Full information - Deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Group Id |
| create_date | [int64](#int64) |  | Date created |
| owner_uid | [int32](#int32) |  | Group owner |
| members | [Member](#dialog.Member) | repeated | Group members. Can be empty when isAsyncMembers enabled. |
| theme | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Group Theme |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Group about |
| is_async_members | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is Members need to be loaded asynchronous. |
| can_view_members | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Can current user view members of the group. Default is true. |
| can_invite_people | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Can current user invite new people. Default is true. |
| is_shared_history | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is history shared among all users. |






<a name="dialog.GroupMemberPermission"/>

### GroupMemberPermission
A struct mapping a group member to their permissions
userId the id of the group member
permissions a list of permissions that user has


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  |  |
| permissions | [GroupAdminPermission](#dialog.GroupAdminPermission) | repeated |  |






<a name="dialog.Member"/>

### Member
Member information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  | User id |
| inviter_uid | [int32](#int32) |  | User inviter id |
| date | [int64](#int64) |  | Adding date |
| is_admin | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Deprecated |
| permissions | [GroupAdminPermission](#dialog.GroupAdminPermission) | repeated | List of member permissions |






<a name="dialog.RequestCreateGroup"/>

### RequestCreateGroup
Creating group chat


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rid | [int64](#int64) |  | Id for query deduplication |
| title | [string](#string) |  |  |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated | members |
| group_type | [GroupType](#dialog.GroupType) |  | group or channel |
| username | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | optional shortname of a group, group will be public if set |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestEditGroupAbout"/>

### RequestEditGroupAbout
Edit Group About


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestEditGroupAvatar"/>

### RequestEditGroupAvatar
Changing group avatar


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| file_location | [FileLocation](#dialog.FileLocation) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestEditGroupTitle"/>

### RequestEditGroupTitle
Changing group title


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| title | [string](#string) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestEditGroupTopic"/>

### RequestEditGroupTopic
Edit group topic


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| topic | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestGetGroupInviteUrl"/>

### RequestGetGroupInviteUrl
Building invite url


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |






<a name="dialog.RequestGetGroupInviteUrlBase"/>

### RequestGetGroupInviteUrlBase
Get group invite url base






<a name="dialog.RequestGetGroupMemberPermissions"/>

### RequestGetGroupMemberPermissions
Fetches the group administration permissions for each of the users from the list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| user_ids | [int32](#int32) | repeated |  |






<a name="dialog.RequestInviteUser"/>

### RequestInviteUser
Inviting user to group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| user | [UserOutPeer](#dialog.UserOutPeer) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestJoinGroup"/>

### RequestJoinGroup
Join group method


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestJoinGroupByPeer"/>

### RequestJoinGroupByPeer
Join group by peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |






<a name="dialog.RequestKickUser"/>

### RequestKickUser
Kicking user from group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| user | [UserOutPeer](#dialog.UserOutPeer) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestLeaveGroup"/>

### RequestLeaveGroup
Leaving group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestLoadFullGroups"/>

### RequestLoadFullGroups
Loading Full Groups - Deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.RequestLoadMembers"/>

### RequestLoadMembers
Loading group members


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| limit | [int32](#int32) |  |  |
| next | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  | cursor |






<a name="dialog.RequestMakeUserAdmin"/>

### RequestMakeUserAdmin
Make user admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| user_peer | [UserOutPeer](#dialog.UserOutPeer) |  |  |
| permissions | [GroupAdminPermission](#dialog.GroupAdminPermission) | repeated |  |






<a name="dialog.RequestMakeUserAdminObsolete"/>

### RequestMakeUserAdminObsolete
[OBSOLETE] Make user admin


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| user_peer | [UserOutPeer](#dialog.UserOutPeer) |  |  |






<a name="dialog.RequestRemoveGroupAvatar"/>

### RequestRemoveGroupAvatar
Removing group avatar


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| rid | [int64](#int64) |  | Id for query deduplication |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated | Optimizations drops some info from response to decrease traffic and latency |






<a name="dialog.RequestRevokeInviteUrl"/>

### RequestRevokeInviteUrl
Revoking invite urls


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |






<a name="dialog.RequestSetGroupShortname"/>

### RequestSetGroupShortname
Sets group short name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| shortname | [string](#string) |  | if shortname was empty, then group will become public |






<a name="dialog.RequestTransferOwnership"/>

### RequestTransferOwnership
Transfer ownership of group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| new_owner | [int32](#int32) |  |  |






<a name="dialog.ResponseCreateGroup"/>

### ResponseCreateGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |
| group | [Group](#dialog.Group) |  | created group |
| users | [User](#dialog.User) | repeated | empty if dropped by optimizations |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated | empty if dropped by optimizations |






<a name="dialog.ResponseEditGroupAvatar"/>

### ResponseEditGroupAvatar



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar | [Avatar](#dialog.Avatar) |  |  |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |
| date | [int64](#int64) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  |  |






<a name="dialog.ResponseGetGroupInviteUrlBase"/>

### ResponseGetGroupInviteUrlBase



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="dialog.ResponseGetGroupMemberPermissions"/>

### ResponseGetGroupMemberPermissions



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permissions | [GroupMemberPermission](#dialog.GroupMemberPermission) | repeated |  |






<a name="dialog.ResponseInviteUrl"/>

### ResponseInviteUrl
Response for invite url methods


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="dialog.ResponseJoinGroup"/>

### ResponseJoinGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#dialog.Group) |  |  |
| users | [User](#dialog.User) | repeated | empty if dropped by optimizations |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated | empty if dropped by optimizations |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.ResponseLoadFullGroups"/>

### ResponseLoadFullGroups
Deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [GroupFull](#dialog.GroupFull) | repeated |  |






<a name="dialog.ResponseLoadMembers"/>

### ResponseLoadMembers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| members | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| next | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ResponseMakeUserAdminObsolete"/>

### ResponseMakeUserAdminObsolete



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| members | [Member](#dialog.Member) | repeated |  |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |






<a name="dialog.UpdateGroupAboutChanged"/>

### UpdateGroupAboutChanged
Update about about changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateGroupAboutChangedObsolete"/>

### UpdateGroupAboutChangedObsolete
Update about group about change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateGroupAvatarChanged"/>

### UpdateGroupAvatarChanged
Update about avatar changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| avatar | [Avatar](#dialog.Avatar) |  |  |






<a name="dialog.UpdateGroupAvatarChangedObsolete"/>

### UpdateGroupAvatarChangedObsolete
Update about group avatar change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| uid | [int32](#int32) |  |  |
| avatar | [Avatar](#dialog.Avatar) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupCanInviteMembersChanged"/>

### UpdateGroupCanInviteMembersChanged
Update about can invite members changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| can_invite_members | [bool](#bool) |  |  |






<a name="dialog.UpdateGroupCanSendMessagesChanged"/>

### UpdateGroupCanSendMessagesChanged
Update about can send messages changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| can_send_messages | [bool](#bool) |  |  |






<a name="dialog.UpdateGroupCanViewMembersChanged"/>

### UpdateGroupCanViewMembersChanged
Update about can view members changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| can_view_members | [bool](#bool) |  |  |






<a name="dialog.UpdateGroupHistoryShared"/>

### UpdateGroupHistoryShared
Update about history shared


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |






<a name="dialog.UpdateGroupInviteObsolete"/>

### UpdateGroupInviteObsolete
Update about inviting current user to group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| rid | [int64](#int64) |  | deprecated |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| invite_uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupMemberAdminChanged"/>

### UpdateGroupMemberAdminChanged
Update about member admin changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| user_id | [int32](#int32) |  |  |
| is_admin | [bool](#bool) |  |  |






<a name="dialog.UpdateGroupMemberChanged"/>

### UpdateGroupMemberChanged
Update about membership changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| is_member | [bool](#bool) |  |  |






<a name="dialog.UpdateGroupMemberDiff"/>

### UpdateGroupMemberDiff
Update about members changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| removed_users | [int32](#int32) | repeated |  |
| added_members | [Member](#dialog.Member) | repeated |  |
| members_count | [int32](#int32) |  |  |






<a name="dialog.UpdateGroupMemberPermissionsChanged"/>

### UpdateGroupMemberPermissionsChanged
Update about the user&#39;s permissions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| user_id | [int32](#int32) |  |  |
| permissions | [GroupAdminPermission](#dialog.GroupAdminPermission) | repeated |  |






<a name="dialog.UpdateGroupMembersBecameAsync"/>

### UpdateGroupMembersBecameAsync
Update about members became async


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |






<a name="dialog.UpdateGroupMembersCountChanged"/>

### UpdateGroupMembersCountChanged
Update about members count changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| members_count | [int32](#int32) |  |  |






<a name="dialog.UpdateGroupMembersUpdateObsolete"/>

### UpdateGroupMembersUpdateObsolete
Silent group members update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| members | [Member](#dialog.Member) | repeated |  |






<a name="dialog.UpdateGroupMembersUpdated"/>

### UpdateGroupMembersUpdated
Update about members updated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| members | [Member](#dialog.Member) | repeated |  |






<a name="dialog.UpdateGroupOwnerChanged"/>

### UpdateGroupOwnerChanged
Update about owner changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| user_id | [int32](#int32) |  | new owner |






<a name="dialog.UpdateGroupShortnameChanged"/>

### UpdateGroupShortnameChanged
Update group short name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| shortname | [string](#string) |  |  |
| uid | [int32](#int32) |  |  |






<a name="dialog.UpdateGroupTitleChanged"/>

### UpdateGroupTitleChanged
Update about title changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| title | [string](#string) |  |  |






<a name="dialog.UpdateGroupTitleChangedObsolete"/>

### UpdateGroupTitleChangedObsolete
Update about group title change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| uid | [int32](#int32) |  |  |
| title | [string](#string) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupTopicChanged"/>

### UpdateGroupTopicChanged
Update about topic changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| topic | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateGroupTopicChangedObsolete"/>

### UpdateGroupTopicChangedObsolete
Update about group topic change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| rid | [int64](#int64) |  | deprecated |
| uid | [int32](#int32) |  |  |
| topic | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupUserInvitedObsolete"/>

### UpdateGroupUserInvitedObsolete
Update about inviting user to group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| rid | [int64](#int64) |  | Id for deduplication |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| uid | [int32](#int32) |  |  |
| inviter_uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupUserKickObsolete"/>

### UpdateGroupUserKickObsolete
Update about kicking user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| uid | [int32](#int32) |  |  |
| kicker_uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.UpdateGroupUserLeaveObsolete"/>

### UpdateGroupUserLeaveObsolete
Update about leaving user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | deprecated |
| uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |





 


<a name="dialog.GroupAdminPermission"/>

### GroupAdminPermission
Possible permissions on a group

| Name | Number | Description |
| ---- | ------ | ----------- |
| GROUPADMINPERMISSION_UNKNOWN | 0 |  |
| GROUPADMINPERMISSION_EDITSHORTNAME | 1 |  |
| GROUPADMINPERMISSION_INVITE | 2 |  |
| GROUPADMINPERMISSION_KICK | 3 |  |
| GROUPADMINPERMISSION_UPDATEINFO | 4 |  |
| GROUPADMINPERMISSION_SETPERMISSIONS | 5 |  |
| GROUPADMINPERMISSION_EDITMESSAGE | 6 |  |
| GROUPADMINPERMISSION_DELETEMESSAGE | 7 |  |
| GROUPADMINPERMISSION_GETINTEGRATIONTOKEN | 8 |  |
| GROUPADMINPERMISSION_SENDMESSAGE | 9 |  |
| GROUPADMINPERMISSION_PINMESSAGE | 10 |  |



<a name="dialog.GroupType"/>

### GroupType


| Name | Number | Description |
| ---- | ------ | ----------- |
| GROUPTYPE_UNKNOWN | 0 |  |
| GROUPTYPE_GROUP | 1 |  |
| GROUPTYPE_CHANNEL | 2 |  |


 

 


<a name="dialog.Groups"/>

### Groups


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LoadFullGroups | [RequestLoadFullGroups](#dialog.RequestLoadFullGroups) | [ResponseLoadFullGroups](#dialog.RequestLoadFullGroups) | deprecated |
| LoadMembers | [RequestLoadMembers](#dialog.RequestLoadMembers) | [ResponseLoadMembers](#dialog.RequestLoadMembers) |  |
| CreateGroup | [RequestCreateGroup](#dialog.RequestCreateGroup) | [ResponseCreateGroup](#dialog.RequestCreateGroup) |  |
| EditGroupTitle | [RequestEditGroupTitle](#dialog.RequestEditGroupTitle) | [ResponseSeqDateMid](#dialog.RequestEditGroupTitle) |  |
| SetGroupShortname | [RequestSetGroupShortname](#dialog.RequestSetGroupShortname) | [ResponseSeq](#dialog.RequestSetGroupShortname) |  |
| EditGroupAvatar | [RequestEditGroupAvatar](#dialog.RequestEditGroupAvatar) | [ResponseEditGroupAvatar](#dialog.RequestEditGroupAvatar) |  |
| RemoveGroupAvatar | [RequestRemoveGroupAvatar](#dialog.RequestRemoveGroupAvatar) | [ResponseSeqDateMid](#dialog.RequestRemoveGroupAvatar) |  |
| EditGroupTopic | [RequestEditGroupTopic](#dialog.RequestEditGroupTopic) | [ResponseSeqDate](#dialog.RequestEditGroupTopic) |  |
| EditGroupAbout | [RequestEditGroupAbout](#dialog.RequestEditGroupAbout) | [ResponseSeqDate](#dialog.RequestEditGroupAbout) |  |
| InviteUser | [RequestInviteUser](#dialog.RequestInviteUser) | [ResponseSeqDateMid](#dialog.RequestInviteUser) |  |
| LeaveGroup | [RequestLeaveGroup](#dialog.RequestLeaveGroup) | [ResponseSeqDateMid](#dialog.RequestLeaveGroup) |  |
| KickUser | [RequestKickUser](#dialog.RequestKickUser) | [ResponseSeqDateMid](#dialog.RequestKickUser) |  |
| MakeUserAdmin | [RequestMakeUserAdmin](#dialog.RequestMakeUserAdmin) | [ResponseSeqDate](#dialog.RequestMakeUserAdmin) |  |
| GetGroupMemberPermissions | [RequestGetGroupMemberPermissions](#dialog.RequestGetGroupMemberPermissions) | [ResponseGetGroupMemberPermissions](#dialog.RequestGetGroupMemberPermissions) |  |
| TransferOwnership | [RequestTransferOwnership](#dialog.RequestTransferOwnership) | [ResponseSeqDate](#dialog.RequestTransferOwnership) |  |
| GetGroupInviteUrl | [RequestGetGroupInviteUrl](#dialog.RequestGetGroupInviteUrl) | [ResponseInviteUrl](#dialog.RequestGetGroupInviteUrl) |  |
| GetGroupInviteUrlBase | [RequestGetGroupInviteUrlBase](#dialog.RequestGetGroupInviteUrlBase) | [ResponseGetGroupInviteUrlBase](#dialog.RequestGetGroupInviteUrlBase) |  |
| RevokeInviteUrl | [RequestRevokeInviteUrl](#dialog.RequestRevokeInviteUrl) | [ResponseInviteUrl](#dialog.RequestRevokeInviteUrl) |  |
| JoinGroup | [RequestJoinGroup](#dialog.RequestJoinGroup) | [ResponseJoinGroup](#dialog.RequestJoinGroup) |  |
| JoinGroupByPeer | [RequestJoinGroupByPeer](#dialog.RequestJoinGroupByPeer) | [ResponseVoid](#dialog.RequestJoinGroupByPeer) |  |
| MakeUserAdminObsolete | [RequestMakeUserAdminObsolete](#dialog.RequestMakeUserAdminObsolete) | [ResponseMakeUserAdminObsolete](#dialog.RequestMakeUserAdminObsolete) |  |

 



<a name="integrations.proto"/>
<p align="right"><a href="#top">Top</a></p>

## integrations.proto



<a name="dialog.RequestGetIntegrationToken"/>

### RequestGetIntegrationToken
Getting current group token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestRevokeIntegrationToken"/>

### RequestRevokeIntegrationToken
Revoke group token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.ResponseIntegrationToken"/>

### ResponseIntegrationToken
Group token response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| url | [string](#string) |  |  |





 

 

 


<a name="dialog.Integrations"/>

### Integrations


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetIntegrationToken | [RequestGetIntegrationToken](#dialog.RequestGetIntegrationToken) | [ResponseIntegrationToken](#dialog.RequestGetIntegrationToken) | Get token for posting to group |
| RevokeIntegrationToken | [RequestRevokeIntegrationToken](#dialog.RequestRevokeIntegrationToken) | [ResponseIntegrationToken](#dialog.RequestRevokeIntegrationToken) | Revoke token |

 



<a name="media_and_files.proto"/>
<p align="right"><a href="#top">Top</a></p>

## media_and_files.proto



<a name="dialog.AudioLocation"/>

### AudioLocation
Audio location


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_location | [FileLocation](#dialog.FileLocation) |  | File location |
| duration | [int32](#int32) |  | audio duration in seconds |
| mime_type | [string](#string) |  | mime type |
| file_size | [int32](#int32) |  | file size |






<a name="dialog.Avatar"/>

### Avatar
Avatar of User or Group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| small_image | [AvatarImage](#dialog.AvatarImage) |  | Optional small image of avatar box in 100x100 |
| large_image | [AvatarImage](#dialog.AvatarImage) |  | Optional large image of avatar box in 200x200 |
| full_image | [AvatarImage](#dialog.AvatarImage) |  | Optional full screen image of avatar |






<a name="dialog.AvatarImage"/>

### AvatarImage
Avatar Image


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_location | [FileLocation](#dialog.FileLocation) |  | Location of file |
| width | [int32](#int32) |  | Width of avatar image |
| height | [int32](#int32) |  | Height of avatar image |
| file_size | [int32](#int32) |  | Size of file |






<a name="dialog.Color"/>

### Color



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rgb | [RgbColor](#dialog.RgbColor) |  |  |
| predefined | [PredefinedColor](#dialog.PredefinedColor) |  |  |






<a name="dialog.FastThumb"/>

### FastThumb
Fast thumb of media messages. Less than 90x90 and compressed by JPEG with low quality


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| w | [int32](#int32) |  | Width of thumb |
| h | [int32](#int32) |  | Height of thump |
| thumb | [bytes](#bytes) |  | compressed image data |






<a name="dialog.FileLocation"/>

### FileLocation
Location of file on server


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_id | [int64](#int64) |  | Unique Id of file |
| access_hash | [int64](#int64) |  | Access hash of file |






<a name="dialog.FileUrlDescription"/>

### FileUrlDescription
File url description


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_id | [int64](#int64) |  | File id of url |
| url | [string](#string) |  | Url for downloading |
| timeout | [int32](#int32) |  | Timeout of url |
| unsigned_url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Unsigned URL (used to honor web caches) |
| unsigned_url_headers | [HTTPHeader](#dialog.HTTPHeader) | repeated | Headers that is required to download files with unsigned url |






<a name="dialog.HTTPHeader"/>

### HTTPHeader
HTTP Header record


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  | HTTP Header name |
| value | [string](#string) |  | HTTP Header value |






<a name="dialog.ImageLocation"/>

### ImageLocation
Image location


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_location | [FileLocation](#dialog.FileLocation) |  | Location of file |
| width | [int32](#int32) |  | Width of avatar image |
| height | [int32](#int32) |  | Height of avatar image |
| file_size | [int32](#int32) |  | Size of file |






<a name="dialog.PredefinedColor"/>

### PredefinedColor
Predefined color


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| color | [Colors](#dialog.Colors) |  | Predefined color value |






<a name="dialog.RequestCommitFileUpload"/>

### RequestCommitFileUpload
Comminting uploaded file to storage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| upload_key | [bytes](#bytes) |  |  |
| file_name | [string](#string) |  |  |






<a name="dialog.RequestGetFileUploadPartUrl"/>

### RequestGetFileUploadPartUrl
Upload file part


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| part_number | [int32](#int32) |  |  |
| part_size | [int32](#int32) |  |  |
| upload_key | [bytes](#bytes) |  |  |






<a name="dialog.RequestGetFileUploadUrl"/>

### RequestGetFileUploadUrl
Requesting pload url


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| expected_size | [int32](#int32) |  |  |






<a name="dialog.RequestGetFileUrl"/>

### RequestGetFileUrl
Requesting file URL for downloading


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file | [FileLocation](#dialog.FileLocation) |  |  |






<a name="dialog.RequestGetFileUrlBuilder"/>

### RequestGetFileUrlBuilder
Get File URL Builder that allows to build file urls from client side


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| supported_signature_algorithms | [string](#string) | repeated |  |






<a name="dialog.RequestGetFileUrls"/>

### RequestGetFileUrls
Requesting multiple fle URL for downloading


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| files | [FileLocation](#dialog.FileLocation) | repeated |  |






<a name="dialog.ResponseCommitFileUpload"/>

### ResponseCommitFileUpload



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uploaded_file_location | [FileLocation](#dialog.FileLocation) |  |  |






<a name="dialog.ResponseGetFileUploadPartUrl"/>

### ResponseGetFileUploadPartUrl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="dialog.ResponseGetFileUploadUrl"/>

### ResponseGetFileUploadUrl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| upload_key | [bytes](#bytes) |  |  |






<a name="dialog.ResponseGetFileUrl"/>

### ResponseGetFileUrl



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |
| timeout | [int32](#int32) |  |  |
| unsigned_url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| unsigned_url_headers | [HTTPHeader](#dialog.HTTPHeader) | repeated |  |






<a name="dialog.ResponseGetFileUrlBuilder"/>

### ResponseGetFileUrlBuilder



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_url | [string](#string) |  |  |
| algo | [string](#string) |  |  |
| seed | [string](#string) |  |  |
| signature_secret | [bytes](#bytes) |  |  |
| timeout | [int32](#int32) |  |  |






<a name="dialog.ResponseGetFileUrls"/>

### ResponseGetFileUrls



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_urls | [FileUrlDescription](#dialog.FileUrlDescription) | repeated |  |






<a name="dialog.RgbColor"/>

### RgbColor
RGB Color


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| rgb | [int32](#int32) |  | RGB color value |





 


<a name="dialog.Colors"/>

### Colors


| Name | Number | Description |
| ---- | ------ | ----------- |
| COLORS_UNKNOWN | 0 |  |
| COLORS_RED | 1 |  |
| COLORS_YELLOW | 2 |  |
| COLORS_GREEN | 3 |  |


 

 


<a name="dialog.MediaAndFiles"/>

### MediaAndFiles


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetFileUrl | [RequestGetFileUrl](#dialog.RequestGetFileUrl) | [ResponseGetFileUrl](#dialog.RequestGetFileUrl) | Get link to file for downloading |
| GetFileUrls | [RequestGetFileUrls](#dialog.RequestGetFileUrls) | [ResponseGetFileUrls](#dialog.RequestGetFileUrls) | Get link to files for downloading |
| GetFileUrlBuilder | [RequestGetFileUrlBuilder](#dialog.RequestGetFileUrlBuilder) | [ResponseGetFileUrlBuilder](#dialog.RequestGetFileUrlBuilder) | Create builder for file url |
| GetFileUploadUrl | [RequestGetFileUploadUrl](#dialog.RequestGetFileUploadUrl) | [ResponseGetFileUploadUrl](#dialog.RequestGetFileUploadUrl) | Get url for uploading |
| CommitFileUpload | [RequestCommitFileUpload](#dialog.RequestCommitFileUpload) | [ResponseCommitFileUpload](#dialog.RequestCommitFileUpload) | Finish uploading a file |
| GetFileUploadPartUrl | [RequestGetFileUploadPartUrl](#dialog.RequestGetFileUploadPartUrl) | [ResponseGetFileUploadPartUrl](#dialog.RequestGetFileUploadPartUrl) | Get url for uploading chunk of file |

 



<a name="messaging.proto"/>
<p align="right"><a href="#top">Top</a></p>

## messaging.proto



<a name="dialog.AudioMedia"/>

### AudioMedia
Audio media


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| audio | [AudioLocation](#dialog.AudioLocation) |  |  |






<a name="dialog.BinaryMessage"/>

### BinaryMessage
Binary Message. Useful for implementing your own content types


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content_tag | [string](#string) |  |  |
| msg | [bytes](#bytes) |  |  |






<a name="dialog.DeletedMessage"/>

### DeletedMessage
Deleted message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_local | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Deleted locally message |






<a name="dialog.Dialog"/>

### Dialog
Conversation from history
peer
unreadCount
sortDate date of conversation for sorting
senderUid Sender of top message (may be zero)
isFavourite Is dialog favourite
rid Random ID of top message (may be zero)
mid Message id
date Date of top message (can&#39;t be zero)
message Content of message
firstUnreadDate Date of first unread message
attributes Optional top message attributes
pinnedMessages Optional pinned messages
historyMessage Optional last messages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  | Peer of conversation |
| unread_count | [int32](#int32) |  | counter of unread messages |
| sort_date | [int64](#int64) |  | deprecated |
| sender_uid | [int32](#int32) |  |  |
| is_favourite | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| date | [int64](#int64) |  | last action date |
| message | [MessageContent](#dialog.MessageContent) |  | last message content |
| state | [MessageState](#dialog.MessageState) |  |  |
| first_unread_date | [int64](#int64) |  |  |
| attributes | [MessageAttributes](#dialog.MessageAttributes) |  |  |
| pinned_messages | [PinnedMessages](#dialog.PinnedMessages) |  |  |
| history_message | [HistoryMessage](#dialog.HistoryMessage) |  | last message in this dialog |






<a name="dialog.DialogGroup"/>

### DialogGroup
Grouped dialog list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| key | [string](#string) |  |  |
| dialogs | [DialogShort](#dialog.DialogShort) | repeated |  |






<a name="dialog.DialogIndex"/>

### DialogIndex
Dialog index from history. Used in client side pagination


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  | The peer for the conversation |
| timestamp | [int64](#int64) |  | The timestamp for the top message |






<a name="dialog.DialogShort"/>

### DialogShort
Short Dialog from grouped conversation list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  | Peer of conversation |
| counter | [int32](#int32) |  | Conversation unread count |
| date | [int64](#int64) |  | Conversation top message date |






<a name="dialog.DocumentEx"/>

### DocumentEx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| photo | [DocumentExPhoto](#dialog.DocumentExPhoto) |  |  |
| video | [DocumentExVideo](#dialog.DocumentExVideo) |  |  |
| voice | [DocumentExVoice](#dialog.DocumentExVoice) |  |  |






<a name="dialog.DocumentExPhoto"/>

### DocumentExPhoto
File photo extension


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| w | [int32](#int32) |  | image width |
| h | [int32](#int32) |  | image height |






<a name="dialog.DocumentExVideo"/>

### DocumentExVideo
File video extension


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| w | [int32](#int32) |  | video width |
| h | [int32](#int32) |  | video height |
| duration | [int32](#int32) |  |  |






<a name="dialog.DocumentExVoice"/>

### DocumentExVoice
File voice extension


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| duration | [int32](#int32) |  |  |






<a name="dialog.DocumentMessage"/>

### DocumentMessage
File message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_id | [int64](#int64) |  |  |
| access_hash | [int64](#int64) |  |  |
| file_size | [int32](#int32) |  |  |
| name | [string](#string) |  |  |
| mime_type | [string](#string) |  |  |
| thumb | [FastThumb](#dialog.FastThumb) |  | optional thumb of file. JPEG less that 90x90 with 60-70 quality. |
| ext | [DocumentEx](#dialog.DocumentEx) |  |  |






<a name="dialog.EmptyMessage"/>

### EmptyMessage
Empty Message






<a name="dialog.HistoryMessage"/>

### HistoryMessage
Message from history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_uid | [int32](#int32) |  |  |
| sender_peer | [OutPeer](#dialog.OutPeer) |  |  |
| host_peer | [OutPeer](#dialog.OutPeer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id generated by server |
| prev_mid | [UUIDValue](#dialog.UUIDValue) |  |  |
| date | [int64](#int64) |  |  |
| message | [MessageContent](#dialog.MessageContent) |  |  |
| state | [MessageState](#dialog.MessageState) |  |  |
| reactions | [MessageReaction](#dialog.MessageReaction) | repeated |  |
| attribute | [MessageAttributes](#dialog.MessageAttributes) |  |  |
| forward | [ReferencedMessages](#dialog.ReferencedMessages) |  |  |
| reply | [ReferencedMessages](#dialog.ReferencedMessages) |  |  |
| edited_at | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |






<a name="dialog.ImageMedia"/>

### ImageMedia
Image media
image image


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| image | [ImageLocation](#dialog.ImageLocation) |  |  |






<a name="dialog.InteractiveMedia"/>

### InteractiveMedia
A text message extension representing an interactive action.
Can be used to add widgets (such as buttons, selects, etc) to messages.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  | identifier of the media action |
| widget | [InteractiveMediaWidget](#dialog.InteractiveMediaWidget) |  | a widget to be shown to user |
| style | [InteractiveMediaStyle](#dialog.InteractiveMediaStyle) |  | a style of the widget, which is interpreted by the client |
| confirm | [InteractiveMediaConfirm](#dialog.InteractiveMediaConfirm) |  | A content of the alert dialog that will be show to user when they perform the media action |






<a name="dialog.InteractiveMediaButton"/>

### InteractiveMediaButton
A &#39;Button&#39; widget


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | A value for this button |
| label | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | A user-visible description of this button |






<a name="dialog.InteractiveMediaConfirm"/>

### InteractiveMediaConfirm
An alert dialog content to show to user


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | the optional alert dialog prompt |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | the optional alert dialog title |
| ok | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | the optional confirm button text |
| dismiss | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | the optional cancel button text |






<a name="dialog.InteractiveMediaGroup"/>

### InteractiveMediaGroup
A group of interactive media actions


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| actions | [InteractiveMedia](#dialog.InteractiveMedia) | repeated | the list of actions |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | an optional title of the group |
| description | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | an optional description of the group |
| translations | [InteractiveMediaTranslationGroup](#dialog.InteractiveMediaTranslationGroup) | repeated | a media content translations |






<a name="dialog.InteractiveMediaSelect"/>

### InteractiveMediaSelect
A select from multiple values widget


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| options | [InteractiveMediaSelectOption](#dialog.InteractiveMediaSelectOption) | repeated | list of values to present to user |
| label | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | A user-visible descripton of this select |
| default_value | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | A value that will be selected by default |






<a name="dialog.InteractiveMediaSelectOption"/>

### InteractiveMediaSelectOption
A row in the select widget


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [string](#string) |  | id of the row |
| label | [string](#string) |  | a user visible text for this row |






<a name="dialog.InteractiveMediaTranslation"/>

### InteractiveMediaTranslation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [string](#string) |  |  |
| value | [string](#string) |  |  |






<a name="dialog.InteractiveMediaTranslationGroup"/>

### InteractiveMediaTranslationGroup



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| language | [string](#string) |  | a translation group language |
| messages | [InteractiveMediaTranslation](#dialog.InteractiveMediaTranslation) | repeated | a list of translation mesages |






<a name="dialog.InteractiveMediaWidget"/>

### InteractiveMediaWidget
Some interactive element inside a message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| interactiveMediaButton | [InteractiveMediaButton](#dialog.InteractiveMediaButton) |  |  |
| interactiveMediaSelect | [InteractiveMediaSelect](#dialog.InteractiveMediaSelect) |  |  |






<a name="dialog.JsonMessage"/>

### JsonMessage
Custom-data JsonMessage


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| raw_json | [string](#string) |  |  |






<a name="dialog.MessageAttributes"/>

### MessageAttributes
Message Attributes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_mentioned | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is mentioned. If set overrides built-in value. |
| is_highlighted | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is message highlighted. Default is false. |
| is_notified | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is notified. If set overrides built-in settings. |
| is_only_for_you | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | If this message is only for you. Default is false |






<a name="dialog.MessageContent"/>

### MessageContent



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| textMessage | [TextMessage](#dialog.TextMessage) |  |  |
| serviceMessage | [ServiceMessage](#dialog.ServiceMessage) |  |  |
| documentMessage | [DocumentMessage](#dialog.DocumentMessage) |  |  |
| jsonMessage | [JsonMessage](#dialog.JsonMessage) |  |  |
| unsupportedMessage | [UnsupportedMessage](#dialog.UnsupportedMessage) |  |  |
| stickerMessage | [StickerMessage](#dialog.StickerMessage) |  |  |
| binaryMessage | [BinaryMessage](#dialog.BinaryMessage) |  |  |
| emptyMessage | [EmptyMessage](#dialog.EmptyMessage) |  |  |
| deletedMessage | [DeletedMessage](#dialog.DeletedMessage) |  |  |






<a name="dialog.MessageMedia"/>

### MessageMedia
Message media


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| webpage | [WebpageMedia](#dialog.WebpageMedia) |  |  |
| image | [ImageMedia](#dialog.ImageMedia) |  |  |
| audio | [AudioMedia](#dialog.AudioMedia) |  |  |
| actions | [InteractiveMediaGroup](#dialog.InteractiveMediaGroup) | repeated |  |






<a name="dialog.MessageReaction"/>

### MessageReaction
Reaction to message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [int32](#int32) | repeated | User&#39;s reaction |
| code | [string](#string) |  | Reaction EMOJI code |






<a name="dialog.ParagraphStyle"/>

### ParagraphStyle
Paragraph style


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| show_paragraph | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Show quote-like paragraph? |
| paragraph_color | [Color](#dialog.Color) |  | Override paragraph color |
| bg_color | [Color](#dialog.Color) |  | Override background color |






<a name="dialog.PinnedMessages"/>

### PinnedMessages
Pinned messages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mids | [UUIDValue](#dialog.UUIDValue) | repeated | Messages ids |






<a name="dialog.QuotedMessage"/>

### QuotedMessage
Quoted Message
messageId
publicGroupId
senderUserId
messageDate
quotedMessageContent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| message_id | [int64](#int64) |  | Message Id if present |
| public_group_id | [int32](#int32) |  | Public Group id if present |
| sender_user_id | [int32](#int32) |  | Sender of message |
| message_date | [int64](#int64) |  | Date of message |
| quoted_message_content | [MessageContent](#dialog.MessageContent) |  | Optional Quoted Message Content. Can be empty if messageId is present and message is in current peer. |






<a name="dialog.ReferencedMessages"/>

### ReferencedMessages
mids Referenced messge ids


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mids | [UUIDValue](#dialog.UUIDValue) | repeated |  |






<a name="dialog.RequestArchiveChat"/>

### RequestArchiveChat
Archiving chat


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestClearChat"/>

### RequestClearChat
Clearing of conversation (without removing dialog from dialogs list)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| last_message_date | [int64](#int64) |  |  |






<a name="dialog.RequestDeleteChat"/>

### RequestDeleteChat
Deleting of conversation (also leave group for group conversations)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| last_message_date | [int64](#int64) |  |  |






<a name="dialog.RequestDeleteMessageObsolete"/>

### RequestDeleteMessageObsolete
Deleting messages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mids | [UUIDValue](#dialog.UUIDValue) | repeated | Message ids |






<a name="dialog.RequestDoInteractiveMediaAction"/>

### RequestDoInteractiveMediaAction
Sends a request to do interactive media message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id of the enclosed message |
| id | [string](#string) |  | Media id |
| value | [string](#string) |  | selected value |






<a name="dialog.RequestFavouriteDialog"/>

### RequestFavouriteDialog
Marking dialog as favourite


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestFetchDialogIndex"/>

### RequestFetchDialogIndex
Fetches dialog index (short info about all user&#39;s dialogs).
Used in client side pagination.






<a name="dialog.RequestHideDialog"/>

### RequestHideDialog
Hide Dialog from grouped list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestLoadArchived"/>

### RequestLoadArchived
Loading archived messages - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_offset | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |
| limit | [int32](#int32) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestLoadDialogs"/>

### RequestLoadDialogs
Loading conversation history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| min_date | [int64](#int64) |  |  |
| limit | [int32](#int32) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |
| filters | [DialogsFilter](#dialog.DialogsFilter) | repeated |  |
| peers_to_load | [Peer](#dialog.Peer) | repeated |  |






<a name="dialog.RequestLoadGroupedDialogs"/>

### RequestLoadGroupedDialogs
Load all dialogs from grouped list - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestLoadHistory"/>

### RequestLoadHistory
Loading history of chat


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| date | [int64](#int64) |  |  |
| load_mode | [ListLoadMode](#dialog.ListLoadMode) |  | forward, backward or both |
| limit | [int32](#int32) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestMessageRead"/>

### RequestMessageRead
Marking plain messages as read


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.RequestMessageReceived"/>

### RequestMessageReceived
Confirmation of plain message receive by device


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.RequestMessageRemoveReaction"/>

### RequestMessageRemoveReaction
Removing Message reaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| code | [string](#string) |  | EMOJI code |






<a name="dialog.RequestMessageSetReaction"/>

### RequestMessageSetReaction
Setting Message reaction


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| code | [string](#string) |  | EMOJI code |






<a name="dialog.RequestNotifyDialogOpened"/>

### RequestNotifyDialogOpened
Notifying about dialog open


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestPinMessage"/>

### RequestPinMessage
Pin message in conversation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| last_pin_date | [int64](#int64) |  |  |






<a name="dialog.RequestSendMessage"/>

### RequestSendMessage
Sending plain message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| rid | [int64](#int64) |  | Id for deduplication |
| message | [MessageContent](#dialog.MessageContent) |  |  |
| is_only_for_user | [int32](#int32) |  | if not empty, then message will be send to this user only |
| forward | [ReferencedMessages](#dialog.ReferencedMessages) |  | If current message forwards some other |
| reply | [ReferencedMessages](#dialog.ReferencedMessages) |  | If current message is a reply on some other |






<a name="dialog.RequestShowDialog"/>

### RequestShowDialog
Show Dialog in grouped list


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestUnfavouriteDialog"/>

### RequestUnfavouriteDialog
Making dialog as unfavourite


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.RequestUnpinMessage"/>

### RequestUnpinMessage
Unpin message in conversation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| last_pin_date | [int64](#int64) |  |  |






<a name="dialog.RequestUpdateMessage"/>

### RequestUpdateMessage
Changing Message content


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| updated_message | [MessageContent](#dialog.MessageContent) |  |  |
| last_edited_at | [int64](#int64) |  | Date from this message when it was changed last time |






<a name="dialog.ResponseDialogsOrder"/>

### ResponseDialogsOrder
Dialogs order response - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| groups | [DialogGroup](#dialog.DialogGroup) | repeated |  |






<a name="dialog.ResponseFetchDialogIndex"/>

### ResponseFetchDialogIndex



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dialog_indices | [DialogIndex](#dialog.DialogIndex) | repeated |  |






<a name="dialog.ResponseLoadArchived"/>

### ResponseLoadArchived
deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [Group](#dialog.Group) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| dialogs | [Dialog](#dialog.Dialog) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| next_offset | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ResponseLoadDialogs"/>

### ResponseLoadDialogs
Contains dialogs and related peers and entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [Group](#dialog.Group) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| dialogs | [Dialog](#dialog.Dialog) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.ResponseLoadGroupedDialogs"/>

### ResponseLoadGroupedDialogs
deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dialogs | [DialogGroup](#dialog.DialogGroup) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| show_archived | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| show_invite | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.ResponseLoadHistory"/>

### ResponseLoadHistory



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| history | [HistoryMessage](#dialog.HistoryMessage) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | counter of unread messages |
| counter_date | [int64](#int64) |  | date, related to this unread counter |






<a name="dialog.ResponseReactionsResponse"/>

### ResponseReactionsResponse
Response for reactions change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |
| reactions | [MessageReaction](#dialog.MessageReaction) | repeated |  |






<a name="dialog.ServiceEx"/>

### ServiceEx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userInvited | [ServiceExUserInvited](#dialog.ServiceExUserInvited) |  |  |
| userJoined | [ServiceExUserJoined](#dialog.ServiceExUserJoined) |  |  |
| userKicked | [ServiceExUserKicked](#dialog.ServiceExUserKicked) |  |  |
| userLeft | [ServiceExUserLeft](#dialog.ServiceExUserLeft) |  |  |
| groupCreated | [ServiceExGroupCreated](#dialog.ServiceExGroupCreated) |  |  |
| changedTitle | [ServiceExChangedTitle](#dialog.ServiceExChangedTitle) |  |  |
| changedTopic | [ServiceExChangedTopic](#dialog.ServiceExChangedTopic) |  |  |
| changedAbout | [ServiceExChangedAbout](#dialog.ServiceExChangedAbout) |  |  |
| changedAvatar | [ServiceExChangedAvatar](#dialog.ServiceExChangedAvatar) |  |  |
| contactRegistered | [ServiceExContactRegistered](#dialog.ServiceExContactRegistered) |  |  |
| phoneMissed | [ServiceExPhoneMissed](#dialog.ServiceExPhoneMissed) |  |  |
| phoneCall | [ServiceExPhoneCall](#dialog.ServiceExPhoneCall) |  |  |
| phoneRejected | [ServiceExPhoneRejected](#dialog.ServiceExPhoneRejected) |  |  |
| chatArchived | [ServiceExChatArchived](#dialog.ServiceExChatArchived) |  |  |
| chatRestored | [ServiceExChatRestored](#dialog.ServiceExChatRestored) |  |  |






<a name="dialog.ServiceExChangedAbout"/>

### ServiceExChangedAbout
Service message on group about change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | New group about |






<a name="dialog.ServiceExChangedAvatar"/>

### ServiceExChangedAvatar
Service message about avatar change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar | [Avatar](#dialog.Avatar) |  | Updated avatar |






<a name="dialog.ServiceExChangedTitle"/>

### ServiceExChangedTitle
Service message about group title change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | New group title |






<a name="dialog.ServiceExChangedTopic"/>

### ServiceExChangedTopic
Service message on group topic change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| topic | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | New group topic |






<a name="dialog.ServiceExChatArchived"/>

### ServiceExChatArchived
Message about chat archived






<a name="dialog.ServiceExChatRestored"/>

### ServiceExChatRestored
Message about chat restored






<a name="dialog.ServiceExContactRegistered"/>

### ServiceExContactRegistered
Service message about user registration


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  | User Id |






<a name="dialog.ServiceExGroupCreated"/>

### ServiceExGroupCreated
Service message about group creating






<a name="dialog.ServiceExPhoneCall"/>

### ServiceExPhoneCall
Update about phone call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| duration | [int32](#int32) |  | Duration of a phone call |






<a name="dialog.ServiceExPhoneMissed"/>

### ServiceExPhoneMissed
Update about missing phone call






<a name="dialog.ServiceExPhoneRejected"/>

### ServiceExPhoneRejected
Update about phone call rejected






<a name="dialog.ServiceExUserInvited"/>

### ServiceExUserInvited
Service message about adding user to group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| invited_uid | [int32](#int32) |  | added user id |






<a name="dialog.ServiceExUserJoined"/>

### ServiceExUserJoined
Service message about user join to group






<a name="dialog.ServiceExUserKicked"/>

### ServiceExUserKicked
Service message about kicking user from group


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| kicked_uid | [int32](#int32) |  |  |






<a name="dialog.ServiceExUserLeft"/>

### ServiceExUserLeft
Service message about user left group






<a name="dialog.ServiceMessage"/>

### ServiceMessage
Service message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| ext | [ServiceEx](#dialog.ServiceEx) |  |  |






<a name="dialog.StickerMessage"/>

### StickerMessage
Sticker message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sticker_id | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | Optional Unique ID of sticker |
| fast_preview | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  | Optional Fast preview of sticker in webp format |
| image_512 | [ImageLocation](#dialog.ImageLocation) |  | Optional 512x512 sticker image in webp format |
| image_256 | [ImageLocation](#dialog.ImageLocation) |  | Optional 256x256 sticker image in webp format |
| sticker_collection_id | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | Optional Collection ID |
| sticker_collection_access_hash | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  | Optional Collection Access Hash |
| emoji | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Sticker emoji |






<a name="dialog.TextCommand"/>

### TextCommand
Text Command Message for bots


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| command | [string](#string) |  | Slash-Command For execution |
| args | [string](#string) |  |  |






<a name="dialog.TextExMarkdown"/>

### TextExMarkdown
Markdown extension


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| markdown | [string](#string) |  | Markdown text |






<a name="dialog.TextMessage"/>

### TextMessage
Text message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |
| mentions | [int32](#int32) | repeated | User mentions in message |
| ext | [TextMessageEx](#dialog.TextMessageEx) |  | Optional bytes of extension |
| media | [MessageMedia](#dialog.MessageMedia) | repeated |  |
| extensions | [Any](#dialog.Any) | repeated |  |






<a name="dialog.TextMessageEx"/>

### TextMessageEx



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| textExMarkdown | [TextExMarkdown](#dialog.TextExMarkdown) |  |  |
| textModernMessage | [TextModernMessage](#dialog.TextModernMessage) |  |  |
| textCommand | [TextCommand](#dialog.TextCommand) |  |  |






<a name="dialog.TextModernAttach"/>

### TextModernAttach
Attaches to message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| title_url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| title_icon | [ImageLocation](#dialog.ImageLocation) |  |  |
| text | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| style | [ParagraphStyle](#dialog.ParagraphStyle) |  |  |
| fields | [TextModernField](#dialog.TextModernField) | repeated |  |






<a name="dialog.TextModernField"/>

### TextModernField
Modern message fields


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |
| value | [string](#string) |  |  |
| is_short | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Is field can be shown in compact way (default is TRUE) |






<a name="dialog.TextModernMessage"/>

### TextModernMessage
Modern text message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| sender_name_override | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| sender_photo_override | [Avatar](#dialog.Avatar) |  |  |
| style | [ParagraphStyle](#dialog.ParagraphStyle) |  |  |
| attaches | [TextModernAttach](#dialog.TextModernAttach) | repeated |  |






<a name="dialog.UnsupportedMessage"/>

### UnsupportedMessage
Explicit type for unsupported message






<a name="dialog.UpdateChatArchive"/>

### UpdateChatArchive
Update about chat archive


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |






<a name="dialog.UpdateChatClear"/>

### UpdateChatClear
Update about chat clear


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| action_date | [int64](#int64) |  |  |






<a name="dialog.UpdateChatDelete"/>

### UpdateChatDelete
Update about chat delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| action_date | [int64](#int64) |  |  |






<a name="dialog.UpdateChatGroupsChanged"/>

### UpdateChatGroupsChanged
Update about chat groups changed. Called only when adding, removing and reordering of grouped dialog.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dialogs | [DialogGroup](#dialog.DialogGroup) | repeated |  |






<a name="dialog.UpdateDialogFavouriteChanged"/>

### UpdateDialogFavouriteChanged
Update about dialog favourite changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| is_favourite | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |






<a name="dialog.UpdateInteractiveMediaEvent"/>

### UpdateInteractiveMediaEvent
The update which will be received when the action is performed. Duplicated on all clients.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id of the message that contains interactive media |
| id | [string](#string) |  | identifier of the media action |
| value | [string](#string) |  | selected value of that action |
| uid | [int32](#int32) |  | who interacted with that media |






<a name="dialog.UpdateMessage"/>

### UpdateMessage
Update about plain message


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| sender_uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| message | [MessageContent](#dialog.MessageContent) |  |  |
| attributes | [MessageAttributes](#dialog.MessageAttributes) |  | attributes to help reasoning about message |
| forward | [ReferencedMessages](#dialog.ReferencedMessages) |  |  |
| reply | [ReferencedMessages](#dialog.ReferencedMessages) |  |  |
| previous_mid | [UUIDValue](#dialog.UUIDValue) |  | Message id of previos message from current conversation |
| counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | counter of unreads message |






<a name="dialog.UpdateMessageContentChanged"/>

### UpdateMessageContentChanged
Update about message change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| message | [MessageContent](#dialog.MessageContent) |  |  |
| edited_at | [int64](#int64) |  |  |






<a name="dialog.UpdateMessageDelete"/>

### UpdateMessageDelete
Update about message delete


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| mids | [UUIDValue](#dialog.UUIDValue) | repeated | Deleted messages |
| counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | counter of unread messages |
| action_date | [int64](#int64) |  | date? related for this unread counter |






<a name="dialog.UpdateMessageRead"/>

### UpdateMessageRead
Update about message read


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| start_date | [int64](#int64) |  | when message was read |
| read_date | [int64](#int64) |  | deprecated |






<a name="dialog.UpdateMessageReadByMe"/>

### UpdateMessageReadByMe
Update about message read by me


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| start_date | [int64](#int64) |  | when message was read |
| unread_counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | counter of unread messages |






<a name="dialog.UpdateMessageReceived"/>

### UpdateMessageReceived
Update about message received


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| start_date | [int64](#int64) |  | when message was receive |
| received_date | [int64](#int64) |  | deprecated |






<a name="dialog.UpdateMessageSent"/>

### UpdateMessageSent
Update about message sent


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| rid | [int64](#int64) |  |  |
| date | [int64](#int64) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | Message id |
| prev_mid | [UUIDValue](#dialog.UUIDValue) |  | Previous message id in current conversation |
| unread_counter | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | counter of unread messages |






<a name="dialog.UpdatePinnedMessagesChanged"/>

### UpdatePinnedMessagesChanged
Update about pinned messages changed in conversation


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| pinned_messages | [PinnedMessages](#dialog.PinnedMessages) |  |  |
| action_date | [int64](#int64) |  |  |






<a name="dialog.UpdateReactionsUpdate"/>

### UpdateReactionsUpdate
Update about reactions change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  | related message id |
| reactions | [MessageReaction](#dialog.MessageReaction) | repeated |  |






<a name="dialog.WebpageMedia"/>

### WebpageMedia
Webpage media


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| description | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| image | [ImageLocation](#dialog.ImageLocation) |  |  |





 


<a name="dialog.DialogsFilter"/>

### DialogsFilter


| Name | Number | Description |
| ---- | ------ | ----------- |
| DIALOGSFILTER_UNKNOWN | 0 |  |
| DIALOGSFILTER_EXCLUDEFAVOURITES | 1 |  |
| DIALOGSFILTER_EXCLUDEARCHIVED | 2 |  |



<a name="dialog.InteractiveMediaStyle"/>

### InteractiveMediaStyle


| Name | Number | Description |
| ---- | ------ | ----------- |
| INTERACTIVEMEDIASTYLE_UNKNOWN | 0 |  |
| INTERACTIVEMEDIASTYLE_DEFAULT | 1 |  |
| INTERACTIVEMEDIASTYLE_PRIMARY | 2 |  |
| INTERACTIVEMEDIASTYLE_DANGER | 3 |  |



<a name="dialog.ListLoadMode"/>

### ListLoadMode


| Name | Number | Description |
| ---- | ------ | ----------- |
| LISTLOADMODE_UNKNOWN | 0 |  |
| LISTLOADMODE_FORWARD | 1 |  |
| LISTLOADMODE_BACKWARD | 2 |  |
| LISTLOADMODE_BOTH | 3 |  |



<a name="dialog.MessageState"/>

### MessageState


| Name | Number | Description |
| ---- | ------ | ----------- |
| MESSAGESTATE_UNKNOWN | 0 |  |
| MESSAGESTATE_SENT | 1 |  |
| MESSAGESTATE_RECEIVED | 2 |  |
| MESSAGESTATE_READ | 3 |  |


 

 


<a name="dialog.Messaging"/>

### Messaging


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DoInteractiveMediaAction | [RequestDoInteractiveMediaAction](#dialog.RequestDoInteractiveMediaAction) | [ResponseVoid](#dialog.RequestDoInteractiveMediaAction) | Interact with a message media (click on button for example) |
| SendMessage | [RequestSendMessage](#dialog.RequestSendMessage) | [ResponseSeqDate](#dialog.RequestSendMessage) |  |
| UpdateMessage | [RequestUpdateMessage](#dialog.RequestUpdateMessage) | [ResponseSeqDate](#dialog.RequestUpdateMessage) |  |
| MessageReceived | [RequestMessageReceived](#dialog.RequestMessageReceived) | [ResponseVoid](#dialog.RequestMessageReceived) | Mark message as received by self |
| MessageRead | [RequestMessageRead](#dialog.RequestMessageRead) | [ResponseVoid](#dialog.RequestMessageRead) | Mark message as read by self |
| DeleteMessageObsolete | [RequestDeleteMessageObsolete](#dialog.RequestDeleteMessageObsolete) | [ResponseSeq](#dialog.RequestDeleteMessageObsolete) | deprecated |
| ClearChat | [RequestClearChat](#dialog.RequestClearChat) | [ResponseSeq](#dialog.RequestClearChat) | Clear chat history |
| DeleteChat | [RequestDeleteChat](#dialog.RequestDeleteChat) | [ResponseSeq](#dialog.RequestDeleteChat) |  |
| ArchiveChat | [RequestArchiveChat](#dialog.RequestArchiveChat) | [ResponseSeq](#dialog.RequestArchiveChat) | deprecated |
| MessageSetReaction | [RequestMessageSetReaction](#dialog.RequestMessageSetReaction) | [ResponseReactionsResponse](#dialog.RequestMessageSetReaction) | Add reaction on message (emoji) |
| MessageRemoveReaction | [RequestMessageRemoveReaction](#dialog.RequestMessageRemoveReaction) | [ResponseReactionsResponse](#dialog.RequestMessageRemoveReaction) |  |
| LoadHistory | [RequestLoadHistory](#dialog.RequestLoadHistory) | [ResponseLoadHistory](#dialog.RequestLoadHistory) | Load conversation history |
| LoadDialogs | [RequestLoadDialogs](#dialog.RequestLoadDialogs) | [ResponseLoadDialogs](#dialog.RequestLoadDialogs) | Load user&#39;s dialogs |
| FetchDialogIndex | [RequestFetchDialogIndex](#dialog.RequestFetchDialogIndex) | [ResponseFetchDialogIndex](#dialog.RequestFetchDialogIndex) | Load short info about all user&#39;s dialogs |
| LoadArchived | [RequestLoadArchived](#dialog.RequestLoadArchived) | [ResponseLoadArchived](#dialog.RequestLoadArchived) | deprecated |
| LoadGroupedDialogs | [RequestLoadGroupedDialogs](#dialog.RequestLoadGroupedDialogs) | [ResponseLoadGroupedDialogs](#dialog.RequestLoadGroupedDialogs) | deprecated |
| HideDialog | [RequestHideDialog](#dialog.RequestHideDialog) | [ResponseDialogsOrder](#dialog.RequestHideDialog) | deprecated |
| ShowDialog | [RequestShowDialog](#dialog.RequestShowDialog) | [ResponseDialogsOrder](#dialog.RequestShowDialog) | deprecated |
| FavouriteDialog | [RequestFavouriteDialog](#dialog.RequestFavouriteDialog) | [ResponseDialogsOrder](#dialog.RequestFavouriteDialog) |  |
| UnfavouriteDialog | [RequestUnfavouriteDialog](#dialog.RequestUnfavouriteDialog) | [ResponseDialogsOrder](#dialog.RequestUnfavouriteDialog) |  |
| NotifyDialogOpened | [RequestNotifyDialogOpened](#dialog.RequestNotifyDialogOpened) | [ResponseVoid](#dialog.RequestNotifyDialogOpened) |  |
| PinMessage | [RequestPinMessage](#dialog.RequestPinMessage) | [ResponseSeqDate](#dialog.RequestPinMessage) |  |
| UnpinMessage | [RequestUnpinMessage](#dialog.RequestUnpinMessage) | [ResponseSeqDate](#dialog.RequestUnpinMessage) |  |

 



<a name="miscellaneous.proto"/>
<p align="right"><a href="#top">Top</a></p>

## miscellaneous.proto



<a name="dialog.Any"/>

### Any
Any
typeUrl type url
data data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type_url | [string](#string) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="dialog.CallsConfig"/>

### CallsConfig
Calls configuration
callsEnabled If true then client should enable calls
videoCallsEnabled If true then client should enable video calls
groupCallsEnabled If true then client should enable group calls
groupCallsMaxMembers Determines how many members may participate in call
rtcpMuxPolicy The RTCP mux policy to use when gathering ICE candidates
emojiSecurityEnabled If true then client should send and render emoji security
screenSharingEnabled If true then client should enable screen sharing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| calls_enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| video_calls_enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| group_calls_enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| group_calls_max_members | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| rtcp_mux_policy | [RtcpMuxPolicy](#dialog.RtcpMuxPolicy) |  |  |
| emoji_security_enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| screen_sharing_enabled | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |






<a name="dialog.Config"/>

### Config
Configuration of system
maxGroupSize Current maximum group size
discover Discover configuration
shareEndpoint Share endpoint
callsConfig Calls configuration
groupInviteConfig Group invite configuration
serverMetaInfo Server meta information (version, etc.)
customProfileSchema Custom profile JSON Schema
servicePeers Service peers supported interpreted by the client


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_group_size | [int32](#int32) |  |  |
| discover | [Discover](#dialog.Discover) |  |  |
| share_endpoint | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| calls_config | [CallsConfig](#dialog.CallsConfig) |  |  |
| group_invite_config | [InvitesConfig](#dialog.InvitesConfig) |  |  |
| server_meta_info | [ServerMetaInfo](#dialog.ServerMetaInfo) |  |  |
| custom_profile_schema | [string](#string) |  |  |
| service_peers | [ServicePeers](#dialog.ServicePeers) |  |  |
| extensions | [Any](#dialog.Any) | repeated |  |






<a name="dialog.Discover"/>

### Discover
Discover description
peers peers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peers | [OutPeer](#dialog.OutPeer) | repeated |  |






<a name="dialog.Extension"/>

### Extension
Extention
id Extension id
data Extension data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |
| data | [bytes](#bytes) |  |  |






<a name="dialog.InvitesConfig"/>

### InvitesConfig
Group invites configuration
baseUrl Base URL part
groupInviteUrlPrefix Prefix for private group invite detection
resolveUrlPrefix Prefix for public groups and users link detection
shareInviteUrl Independent URL for external app sharing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| base_url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| group_invite_url_prefix | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| resolve_url_prefix | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| share_invite_url | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RecursiveMapValue"/>

### RecursiveMapValue



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| items | [RecursiveMapValue.Item](#dialog.RecursiveMapValue.Item) | repeated |  |






<a name="dialog.RecursiveMapValue.Array"/>

### RecursiveMapValue.Array



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| values | [RecursiveMapValue.Value](#dialog.RecursiveMapValue.Value) | repeated |  |






<a name="dialog.RecursiveMapValue.Item"/>

### RecursiveMapValue.Item



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| key | [string](#string) |  |  |
| value | [RecursiveMapValue.Value](#dialog.RecursiveMapValue.Value) |  |  |






<a name="dialog.RecursiveMapValue.Value"/>

### RecursiveMapValue.Value



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| d | [google.protobuf.DoubleValue](#google.protobuf.DoubleValue) |  |  |
| i32 | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| i64 | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |
| str | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| rec | [RecursiveMapValue](#dialog.RecursiveMapValue) |  |  |
| array_rec | [RecursiveMapValue.Array](#dialog.RecursiveMapValue.Array) |  |  |






<a name="dialog.ResponseBool"/>

### ResponseBool
Boolean response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| value | [bool](#bool) |  |  |






<a name="dialog.ResponseSeq"/>

### ResponseSeq
Sequence response. Methods that return this value must process response in particular order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| date | [int64](#int64) |  |  |






<a name="dialog.ResponseSeqDate"/>

### ResponseSeqDate
Sequence response with date. Methods that return this value must process response in particular order


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| date | [int64](#int64) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  |  |






<a name="dialog.ResponseSeqDateMid"/>

### ResponseSeqDateMid
Response with seq, date and messageId


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| date | [int64](#int64) |  |  |
| mid | [UUIDValue](#dialog.UUIDValue) |  |  |






<a name="dialog.ResponseVoid"/>

### ResponseVoid
Empty response






<a name="dialog.ServerMetaInfo"/>

### ServerMetaInfo
Some info about the servr
releaseDate Server release date in milliseconds
apiVersion Server api version


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| release_date | [int64](#int64) |  |  |
| api_version | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |






<a name="dialog.ServicePeers"/>

### ServicePeers
Service peers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| security | [OutPeer](#dialog.OutPeer) |  |  |
| support | [OutPeer](#dialog.OutPeer) |  |  |
| stash | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.UpdateConfig"/>

### UpdateConfig
Update about config change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#dialog.Config) |  |  |
| config_hash | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |





 


<a name="dialog.RtcpMuxPolicy"/>

### RtcpMuxPolicy


| Name | Number | Description |
| ---- | ------ | ----------- |
| RTCPMUXPOLICY_UNKNOWN | 0 |  |
| RTCPMUXPOLICY_NEGOTIATE | 1 |  |
| RTCPMUXPOLICY_REQUIRE | 2 |  |



<a name="dialog.UpdateOptimization"/>

### UpdateOptimization


| Name | Number | Description |
| ---- | ------ | ----------- |
| UPDATEOPTIMIZATION_UNKNOWN | 0 |  |
| UPDATEOPTIMIZATION_NONE | 1 |  |
| UPDATEOPTIMIZATION_STRIP_ENTITIES | 2 |  |
| UPDATEOPTIMIZATION_ENABLE_COMBINED | 3 |  |
| UPDATEOPTIMIZATION_FASTER_MESSAGES | 4 |  |
| UPDATEOPTIMIZATION_STRIP_COUNTERS | 5 |  |
| UPDATEOPTIMIZATION_COMPACT_USERS | 6 |  |
| UPDATEOPTIMIZATION_GROUPS_V2 | 7 |  |
| UPDATEOPTIMIZATION_STRIP_ENTITIES_V2 | 8 |  |


 

 

 



<a name="obsolete.proto"/>
<p align="right"><a href="#top">Top</a></p>

## obsolete.proto



<a name="dialog.ObsoleteGetDifferenceCommand"/>

### ObsoleteGetDifferenceCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| state | [bytes](#bytes) |  |  |
| configHash | [int64](#int64) |  |  |






<a name="dialog.ObsoleteOutPeer"/>

### ObsoleteOutPeer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [ObsoletePeer](#dialog.ObsoletePeer) |  |  |
| accessHash | [int64](#int64) |  |  |






<a name="dialog.ObsoletePeer"/>

### ObsoletePeer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ObsoletePeer.ObsoletePeerType](#dialog.ObsoletePeer.ObsoletePeerType) |  |  |
| id | [int32](#int32) |  |  |
| strId | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| access_hash | [int64](#int64) |  |  |






<a name="dialog.ObsoletePeersList"/>

### ObsoletePeersList



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peers | [ObsoletePeer](#dialog.ObsoletePeer) | repeated |  |






<a name="dialog.ObsoleteSeqUpdateBox"/>

### ObsoleteSeqUpdateBox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| obsoleteUpdate | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ObsoleteServiceUpdate"/>

### ObsoleteServiceUpdate



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| obsoleteUpdate | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox"/>

### ObsoleteWeakUpdateBox



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [int64](#int64) |  |  |
| typing | [ObsoleteWeakUpdateBox.ObsoleteUpdateTyping](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateTyping) |  |  |
| userLastSeen | [ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen) |  |  |
| groupOnline | [ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline) |  |  |
| busMessage | [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage) |  |  |
| busDeviceConnected | [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected) |  |  |
| busDeviceDisconnected | [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected) |  |  |
| busDisposed | [ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed) |  |  |
| forceReload | [ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState) |  |  |
| incomingCall | [ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall) |  |  |
| callHandled | [ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled) |  |  |
| callDisposed | [ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| callId | [int64](#int64) |  |  |
| attemptIndex | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| reason | [ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed.ObsoleteDisposalReason](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallDisposed.ObsoleteDisposalReason) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateCallHandled



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| callId | [int64](#int64) |  |  |
| attemptIndex | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceConnected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| busId | [string](#string) |  |  |
| userId | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| deviceId | [int64](#int64) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDeviceDisconnected



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| busId | [string](#string) |  |  |
| userId | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| deviceId | [int64](#int64) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusDisposed



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| busId | [string](#string) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateEventBusMessage



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| busId | [string](#string) |  |  |
| senderId | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| senderDeviceId | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |
| message | [bytes](#bytes) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| fields | [ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState.ObsoleteForceReloadField](#dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateForceReloadState.ObsoleteForceReloadField) | repeated |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateGroupOnline



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groupId | [int32](#int32) |  |  |
| count | [int32](#int32) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateIncomingCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| callId | [int64](#int64) |  |  |
| busId | [string](#string) |  |  |
| peer | [ObsoletePeer](#dialog.ObsoletePeer) |  |  |
| displayName | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| attemptIndex | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| outPeer | [ObsoleteOutPeer](#dialog.ObsoleteOutPeer) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateTyping"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateTyping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [ObsoletePeer](#dialog.ObsoletePeer) |  |  |
| userId | [int32](#int32) |  |  |
| type | [ObsoleteTypingType](#dialog.ObsoleteTypingType) |  |  |
| isTyping | [bool](#bool) |  |  |






<a name="dialog.ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen"/>

### ObsoleteWeakUpdateBox.ObsoleteUpdateUserLastSeen



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| userId | [int32](#int32) |  |  |
| epochMillis | [int64](#int64) |  |  |
| deviceType | [int32](#int32) |  |  |
| deviceCategory | [string](#string) |  |  |
| isOnline | [bool](#bool) |  |  |






<a name="dialog.ObsoleteWeakUpdateCommand"/>

### ObsoleteWeakUpdateCommand



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| subscribeToOnlines | [ObsoletePeersList](#dialog.ObsoletePeersList) |  |  |
| unsubscribeFromOnlines | [ObsoletePeersList](#dialog.ObsoletePeersList) |  |  |
| dropOnlineSubscriptions | [google.protobuf.Empty](#google.protobuf.Empty) |  |  |
| myTyping | [ObsoleteWeakUpdateCommand.ObsoleteMyTyping](#dialog.ObsoleteWeakUpdateCommand.ObsoleteMyTyping) |  |  |
| subscribeToTypings | [ObsoletePeersList](#dialog.ObsoletePeersList) |  |  |
| unsubscribeFromTypings | [ObsoletePeersList](#dialog.ObsoletePeersList) |  |  |
| dropTypingSubscriptions | [google.protobuf.Empty](#google.protobuf.Empty) |  |  |
| subscribeToEventBusUpdates | [google.protobuf.Empty](#google.protobuf.Empty) |  |  |
| myOnline | [ObsoleteWeakUpdateCommand.ObsoleteMyOnline](#dialog.ObsoleteWeakUpdateCommand.ObsoleteMyOnline) |  |  |






<a name="dialog.ObsoleteWeakUpdateCommand.ObsoleteMyOnline"/>

### ObsoleteWeakUpdateCommand.ObsoleteMyOnline



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| online | [bool](#bool) |  |  |






<a name="dialog.ObsoleteWeakUpdateCommand.ObsoleteMyTyping"/>

### ObsoleteWeakUpdateCommand.ObsoleteMyTyping



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [ObsoletePeer](#dialog.ObsoletePeer) |  |  |
| type | [ObsoleteTypingType](#dialog.ObsoleteTypingType) |  |  |
| start | [bool](#bool) |  |  |





 


<a name="dialog.ObsoletePeer.ObsoletePeerType"/>

### ObsoletePeer.ObsoletePeerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| OBSOLETE_PEERTYPE_UNKNOWN | 0 |  |
| OBSOLETE_PEERTYPE_PRIVATE | 1 |  |
| OBSOLETE_PEERTYPE_GROUP | 2 |  |
| OBSOLETE_PEERTYPE_SIP | 4 | PEERTYPE_ENCRYPTED_PRIVATE = 3; |



<a name="dialog.ObsoleteTypingType"/>

### ObsoleteTypingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| OBSOLETE_TYPINGTYPE_UNKNOWN | 0 |  |
| OBSOLETE_TYPINGTYPE_TEXT | 1 |  |


 

 


<a name="dialog.Obsolete"/>

### Obsolete
deprecated

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Obsolete | [.google.protobuf.BytesValue](#google.protobuf.BytesValue) | [.google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |
| SeqUpdates | [.google.protobuf.Empty](#google.protobuf.Empty) | [ObsoleteSeqUpdateBox](#google.protobuf.Empty) |  |
| WeakUpdates | [ObsoleteWeakUpdateCommand](#dialog.ObsoleteWeakUpdateCommand) | [ObsoleteWeakUpdateBox](#dialog.ObsoleteWeakUpdateCommand) |  |

 



<a name="peers.proto"/>
<p align="right"><a href="#top">Top</a></p>

## peers.proto



<a name="dialog.GroupOutPeer"/>

### GroupOutPeer
Group&#39;s out peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| access_hash | [int64](#int64) |  |  |






<a name="dialog.OutPeer"/>

### OutPeer
Out peer with access hash


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [PeerType](#dialog.PeerType) |  |  |
| id | [int32](#int32) |  |  |
| access_hash | [int64](#int64) |  |  |
| str_id | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.Peer"/>

### Peer
Peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [PeerType](#dialog.PeerType) |  |  |
| id | [int32](#int32) |  |  |
| str_id | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UserOutPeer"/>

### UserOutPeer
User&#39;s out peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  | User&#39;s id |
| access_hash | [int64](#int64) |  |  |





 


<a name="dialog.PeerType"/>

### PeerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| PEERTYPE_UNKNOWN | 0 |  |
| PEERTYPE_PRIVATE | 1 |  |
| PEERTYPE_GROUP | 2 |  |
| PEERTYPE_ENCRYPTEDPRIVATE | 3 |  |
| PEERTYPE_SIP | 4 |  |


 

 

 



<a name="privacy.proto"/>
<p align="right"><a href="#top">Top</a></p>

## privacy.proto



<a name="dialog.RequestBlockUser"/>

### RequestBlockUser
Block User


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [UserOutPeer](#dialog.UserOutPeer) |  |  |






<a name="dialog.RequestLoadBlockedUsers"/>

### RequestLoadBlockedUsers
Load Blocked Users






<a name="dialog.RequestUnblockUser"/>

### RequestUnblockUser
Unblock User


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [UserOutPeer](#dialog.UserOutPeer) |  |  |






<a name="dialog.ResponseLoadBlockedUsers"/>

### ResponseLoadBlockedUsers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.UpdateUserBlocked"/>

### UpdateUserBlocked
Update about User Blocked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |






<a name="dialog.UpdateUserUnblocked"/>

### UpdateUserUnblocked
Update about User Unblocked


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |





 

 

 


<a name="dialog.Privacy"/>

### Privacy


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| BlockUser | [RequestBlockUser](#dialog.RequestBlockUser) | [ResponseSeq](#dialog.RequestBlockUser) |  |
| UnblockUser | [RequestUnblockUser](#dialog.RequestUnblockUser) | [ResponseSeq](#dialog.RequestUnblockUser) |  |
| LoadBlockedUsers | [RequestLoadBlockedUsers](#dialog.RequestLoadBlockedUsers) | [ResponseLoadBlockedUsers](#dialog.RequestLoadBlockedUsers) |  |

 



<a name="profile.proto"/>
<p align="right"><a href="#top">Top</a></p>

## profile.proto



<a name="dialog.RequestChangeUserStatus"/>

### RequestChangeUserStatus
Changing user&#39;s status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| status | [UserStatus](#dialog.UserStatus) |  |  |






<a name="dialog.RequestCheckNickName"/>

### RequestCheckNickName
Checking availability of nickname


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [string](#string) |  |  |






<a name="dialog.RequestEditAbout"/>

### RequestEditAbout
Changing about information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestEditAvatar"/>

### RequestEditAvatar
Changing account&#39;s avatar


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| file_location | [FileLocation](#dialog.FileLocation) |  |  |






<a name="dialog.RequestEditCustomProfile"/>

### RequestEditCustomProfile
Chaning user custom profile based on scheme


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| custom_profile | [string](#string) |  |  |






<a name="dialog.RequestEditMyPreferredLanguages"/>

### RequestEditMyPreferredLanguages
Changing preffered languages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| preferred_languages | [string](#string) | repeated |  |






<a name="dialog.RequestEditMyTimeZone"/>

### RequestEditMyTimeZone
Updating user&#39;s time zone


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tz | [string](#string) |  |  |






<a name="dialog.RequestEditName"/>

### RequestEditName
Changing account&#39;s name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="dialog.RequestEditNickName"/>

### RequestEditNickName
Changing account&#39;s nickname


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| nickname | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestEditSex"/>

### RequestEditSex
Changing user&#39;s sex


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sex | [Sex](#dialog.Sex) |  |  |






<a name="dialog.RequestRemoveAvatar"/>

### RequestRemoveAvatar
Removing account&#39;s avatar






<a name="dialog.ResponseEditAvatar"/>

### ResponseEditAvatar



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| avatar | [Avatar](#dialog.Avatar) |  |  |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |





 

 

 


<a name="dialog.Profile"/>

### Profile


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| EditName | [RequestEditName](#dialog.RequestEditName) | [ResponseSeq](#dialog.RequestEditName) |  |
| EditNickName | [RequestEditNickName](#dialog.RequestEditNickName) | [ResponseSeq](#dialog.RequestEditNickName) |  |
| CheckNickName | [RequestCheckNickName](#dialog.RequestCheckNickName) | [ResponseBool](#dialog.RequestCheckNickName) |  |
| EditAbout | [RequestEditAbout](#dialog.RequestEditAbout) | [ResponseSeq](#dialog.RequestEditAbout) |  |
| EditAvatar | [RequestEditAvatar](#dialog.RequestEditAvatar) | [ResponseEditAvatar](#dialog.RequestEditAvatar) |  |
| RemoveAvatar | [RequestRemoveAvatar](#dialog.RequestRemoveAvatar) | [ResponseSeq](#dialog.RequestRemoveAvatar) |  |
| EditMyTimeZone | [RequestEditMyTimeZone](#dialog.RequestEditMyTimeZone) | [ResponseSeq](#dialog.RequestEditMyTimeZone) |  |
| EditMyPreferredLanguages | [RequestEditMyPreferredLanguages](#dialog.RequestEditMyPreferredLanguages) | [ResponseSeq](#dialog.RequestEditMyPreferredLanguages) |  |
| EditSex | [RequestEditSex](#dialog.RequestEditSex) | [ResponseSeq](#dialog.RequestEditSex) |  |
| EditCustomProfile | [RequestEditCustomProfile](#dialog.RequestEditCustomProfile) | [ResponseSeq](#dialog.RequestEditCustomProfile) |  |
| ChangeUserStatus | [RequestChangeUserStatus](#dialog.RequestChangeUserStatus) | [ResponseSeq](#dialog.RequestChangeUserStatus) |  |

 



<a name="push.proto"/>
<p align="right"><a href="#top">Top</a></p>

## push.proto



<a name="dialog.RequestRegisterApplePush"/>

### RequestRegisterApplePush
Registering apple push on server


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apns_key | [int32](#int32) |  |  |
| token | [string](#string) |  |  |






<a name="dialog.RequestRegisterApplePushKit"/>

### RequestRegisterApplePushKit
Registration of a new Apple&#39;s PushKit tokens


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| apns_key | [int32](#int32) |  |  |
| token | [string](#string) |  |  |






<a name="dialog.RequestRegisterApplePushToken"/>

### RequestRegisterApplePushToken
Registering Apple Push Token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| bundle_id | [string](#string) |  |  |
| token | [string](#string) |  |  |






<a name="dialog.RequestRegisterGooglePush"/>

### RequestRegisterGooglePush
Registering push token on server


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| project_id | [int64](#int64) |  |  |
| token | [string](#string) |  |  |






<a name="dialog.RequestUnregisterApplePush"/>

### RequestUnregisterApplePush
Unregistering Apple Push


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="dialog.RequestUnregisterApplePushKit"/>

### RequestUnregisterApplePushKit
Unregistering Apple Push Kit token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="dialog.RequestUnregisterApplePushToken"/>

### RequestUnregisterApplePushToken
Unregister Apple Push token


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |






<a name="dialog.RequestUnregisterGooglePush"/>

### RequestUnregisterGooglePush
Unregistering Google Push


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |





 

 

 


<a name="dialog.Push"/>

### Push


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RegisterGooglePush | [RequestRegisterGooglePush](#dialog.RequestRegisterGooglePush) | [ResponseVoid](#dialog.RequestRegisterGooglePush) |  |
| UnregisterGooglePush | [RequestUnregisterGooglePush](#dialog.RequestUnregisterGooglePush) | [ResponseVoid](#dialog.RequestUnregisterGooglePush) |  |
| RegisterApplePush | [RequestRegisterApplePush](#dialog.RequestRegisterApplePush) | [ResponseVoid](#dialog.RequestRegisterApplePush) |  |
| UnregisterApplePush | [RequestUnregisterApplePush](#dialog.RequestUnregisterApplePush) | [ResponseVoid](#dialog.RequestUnregisterApplePush) |  |
| RegisterApplePushKit | [RequestRegisterApplePushKit](#dialog.RequestRegisterApplePushKit) | [ResponseVoid](#dialog.RequestRegisterApplePushKit) |  |
| UnregisterApplePushKit | [RequestUnregisterApplePushKit](#dialog.RequestUnregisterApplePushKit) | [ResponseVoid](#dialog.RequestUnregisterApplePushKit) |  |
| RegisterApplePushToken | [RequestRegisterApplePushToken](#dialog.RequestRegisterApplePushToken) | [ResponseVoid](#dialog.RequestRegisterApplePushToken) |  |
| UnregisterApplePushToken | [RequestUnregisterApplePushToken](#dialog.RequestUnregisterApplePushToken) | [ResponseVoid](#dialog.RequestUnregisterApplePushToken) |  |

 



<a name="raw_api.proto"/>
<p align="right"><a href="#top">Top</a></p>

## raw_api.proto



<a name="dialog.RequestRawRequest"/>

### RequestRawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| service | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| body | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.ResponseRawRequest"/>

### ResponseRawRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| body | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |





 

 

 


<a name="dialog.RawAPI"/>

### RawAPI


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| RawRequest | [RequestRawRequest](#dialog.RequestRawRequest) | [ResponseRawRequest](#dialog.RequestRawRequest) |  |

 



<a name="registration.proto"/>
<p align="right"><a href="#top">Top</a></p>

## registration.proto



<a name="dialog.RegisterDeprecatedDeviceRequest"/>

### RegisterDeprecatedDeviceRequest







<a name="dialog.RequestExchangeAuthIdForToken"/>

### RequestExchangeAuthIdForToken



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| auth_id | [int64](#int64) |  |  |
| signature | [bytes](#bytes) |  |  |






<a name="dialog.RequestRegisterDevice"/>

### RequestRegisterDevice



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| client_pk | [bytes](#bytes) |  |  |
| app_id | [int32](#int32) |  |  |
| app_title | [string](#string) |  |  |
| device_title | [string](#string) |  |  |






<a name="dialog.ResponseDeviceRequest"/>

### ResponseDeviceRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server_pk | [bytes](#bytes) |  |  |
| auth_id | [int64](#int64) |  |  |
| token | [string](#string) |  |  |





 

 

 


<a name="dialog.Registration"/>

### Registration


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ExchangeAuthIdForToken | [RequestExchangeAuthIdForToken](#dialog.RequestExchangeAuthIdForToken) | [ResponseDeviceRequest](#dialog.RequestExchangeAuthIdForToken) |  |
| RegisterDevice | [RequestRegisterDevice](#dialog.RequestRegisterDevice) | [ResponseDeviceRequest](#dialog.RequestRegisterDevice) | register device to get auth token |
| RegisterDeprecatedDevice | [RegisterDeprecatedDeviceRequest](#dialog.RegisterDeprecatedDeviceRequest) | [ResponseDeviceRequest](#dialog.RegisterDeprecatedDeviceRequest) | deprecated |

 



<a name="search.proto"/>
<p align="right"><a href="#top">Top</a></p>

## search.proto



<a name="dialog.MessageSearchItem"/>

### MessageSearchItem
Message Search result container


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| result | [MessageSearchResult](#dialog.MessageSearchResult) |  |  |






<a name="dialog.MessageSearchResult"/>

### MessageSearchResult
Message container


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| rid | [int64](#int64) |  | deprecated |
| date | [int64](#int64) |  |  |
| sender_id | [int32](#int32) |  |  |
| content | [MessageContent](#dialog.MessageContent) |  |  |






<a name="dialog.PeerSearchResult"/>

### PeerSearchResult
Peer search result


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| title | [string](#string) |  |  |
| shortname | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| description | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| members_count | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| date_created | [int64](#int64) |  | Group Creation Date |
| creator | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| is_public | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| is_joined | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Are you joined? |






<a name="dialog.RequestMessageSearch"/>

### RequestMessageSearch
Performing message search


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [SearchCondition](#dialog.SearchCondition) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestMessageSearchMore"/>

### RequestMessageSearchMore
Performing message search paging


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| load_more_state | [bytes](#bytes) |  | Cursor |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestPeerSearch"/>

### RequestPeerSearch
Performing peer search


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [SearchCondition](#dialog.SearchCondition) | repeated |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestResolvePeer"/>

### RequestResolvePeer
Resolve peer by shortname


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| shortname | [string](#string) |  |  |






<a name="dialog.RequestSimpleSearch"/>

### RequestSimpleSearch



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| criteria | [SimpleSearchCondition](#dialog.SimpleSearchCondition) | repeated |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestSimpleSearchMore"/>

### RequestSimpleSearchMore



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| load_more_state | [bytes](#bytes) |  | Cursor |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.ResponseMessageSearchResponse"/>

### ResponseMessageSearchResponse
Search Result with related peers and entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| search_results | [MessageSearchItem](#dialog.MessageSearchItem) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| load_more_state | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |
| user_out_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| group_out_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.ResponsePeerSearch"/>

### ResponsePeerSearch
Response with related peers and entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| search_results | [PeerSearchResult](#dialog.PeerSearchResult) | repeated |  |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.ResponseResolvePeer"/>

### ResponseResolvePeer



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.SearchAndCondition"/>

### SearchAndCondition
Search AND condion


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| and_query | [SearchCondition](#dialog.SearchCondition) | repeated | &#34;And&#34; query |






<a name="dialog.SearchCondition"/>

### SearchCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| searchPeerTypeCondition | [SearchPeerTypeCondition](#dialog.SearchPeerTypeCondition) |  |  |
| searchPieceText | [SearchPieceText](#dialog.SearchPieceText) |  |  |
| searchAndCondition | [SearchAndCondition](#dialog.SearchAndCondition) |  |  |
| searchOrCondition | [SearchOrCondition](#dialog.SearchOrCondition) |  |  |
| searchPeerCondition | [SearchPeerCondition](#dialog.SearchPeerCondition) |  |  |
| searchPeerContentType | [SearchPeerContentType](#dialog.SearchPeerContentType) |  |  |
| searchSenderIdConfition | [SearchSenderIdConfition](#dialog.SearchSenderIdConfition) |  |  |






<a name="dialog.SearchOrCondition"/>

### SearchOrCondition
Search OR condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| or_query | [SearchCondition](#dialog.SearchCondition) | repeated | &#34;Or&#34; query |






<a name="dialog.SearchPeerCondition"/>

### SearchPeerCondition
Serch Peer condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |






<a name="dialog.SearchPeerContentType"/>

### SearchPeerContentType
Search content type condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| content_type | [SearchContentType](#dialog.SearchContentType) |  |  |






<a name="dialog.SearchPeerTypeCondition"/>

### SearchPeerTypeCondition
Search peer type condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer_type | [SearchPeerType](#dialog.SearchPeerType) |  |  |






<a name="dialog.SearchPieceText"/>

### SearchPieceText
Search peer name condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query | [string](#string) |  |  |






<a name="dialog.SearchSenderIdConfition"/>

### SearchSenderIdConfition
Searching sender uid condition


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sender_id | [int32](#int32) |  |  |






<a name="dialog.SimpleContactSearchCondition"/>

### SimpleContactSearchCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| text | [string](#string) |  |  |






<a name="dialog.SimpleMessageSearchCondition"/>

### SimpleMessageSearchCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  | where to search |
| text | [string](#string) |  | search term |
| content_type | [SearchContentType](#dialog.SearchContentType) |  | content message type to search |






<a name="dialog.SimplePeerSearchCondition"/>

### SimplePeerSearchCondition
Search among contacts/groups/users


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer_type | [SearchPeerType](#dialog.SearchPeerType) |  |  |
| text | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.SimpleSearchCondition"/>

### SimpleSearchCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| contact | [SimpleContactSearchCondition](#dialog.SimpleContactSearchCondition) |  |  |
| message | [SimpleMessageSearchCondition](#dialog.SimpleMessageSearchCondition) |  |  |
| peer | [SimplePeerSearchCondition](#dialog.SimplePeerSearchCondition) |  |  |
| userProfile | [SimpleUserProfileSearchCondition](#dialog.SimpleUserProfileSearchCondition) |  |  |






<a name="dialog.SimpleUserProfileSearchCondition"/>

### SimpleUserProfileSearchCondition



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| query_string | [string](#string) |  |  |






<a name="dialog.criterion"/>

### criterion






 


<a name="dialog.SearchContentType"/>

### SearchContentType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEARCHCONTENTTYPE_UNKNOWN | 0 |  |
| SEARCHCONTENTTYPE_ANY | 1 |  |
| SEARCHCONTENTTYPE_TEXT | 2 |  |
| SEARCHCONTENTTYPE_LINKS | 3 |  |
| SEARCHCONTENTTYPE_DOCUMENTS | 4 |  |
| SEARCHCONTENTTYPE_PHOTOS | 5 |  |



<a name="dialog.SearchPeerType"/>

### SearchPeerType


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEARCHPEERTYPE_UNKNOWN | 0 |  |
| SEARCHPEERTYPE_GROUPS | 1 |  |
| SEARCHPEERTYPE_CONTACTS | 2 |  |
| SEARCHPEERTYPE_PUBLIC | 3 |  |


 

 


<a name="dialog.Search"/>

### Search


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| PeerSearch | [RequestPeerSearch](#dialog.RequestPeerSearch) | [ResponsePeerSearch](#dialog.RequestPeerSearch) | Search among groups/users/contacts |
| ResolvePeer | [RequestResolvePeer](#dialog.RequestResolvePeer) | [ResponseResolvePeer](#dialog.RequestResolvePeer) |  |
| MessageSearch | [RequestMessageSearch](#dialog.RequestMessageSearch) | [ResponseMessageSearchResponse](#dialog.RequestMessageSearch) | Search by messages |
| MessageSearchMore | [RequestMessageSearchMore](#dialog.RequestMessageSearchMore) | [ResponseMessageSearchResponse](#dialog.RequestMessageSearchMore) |  |
| SimpleSearch | [RequestSimpleSearch](#dialog.RequestSimpleSearch) | [ResponseMessageSearchResponse](#dialog.RequestSimpleSearch) | Custom search by conditions |
| SimpleSearchMore | [RequestSimpleSearchMore](#dialog.RequestSimpleSearchMore) | [ResponseMessageSearchResponse](#dialog.RequestSimpleSearchMore) |  |

 



<a name="sequence_and_updates.proto"/>
<p align="right"><a href="#top">Top</a></p>

## sequence_and_updates.proto



<a name="dialog.GroupMembersSubset"/>

### GroupMembersSubset
Represents subset of group members


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_peer | [GroupOutPeer](#dialog.GroupOutPeer) |  |  |
| member_ids | [int32](#int32) | repeated |  |






<a name="dialog.RequestGetDialogsDifference"/>

### RequestGetDialogsDifference
Getting difference of dialogs


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| clock | [int64](#int64) |  | max of the last action dates among all user&#39;s dialogs |






<a name="dialog.RequestGetDifference"/>

### RequestGetDifference
Getting difference of sequence


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |
| config_hash | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |






<a name="dialog.RequestGetReferencedEntitites"/>

### RequestGetReferencedEntitites
Loading referenced entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated | users needed |
| mids | [UUIDValue](#dialog.UUIDValue) | repeated | messages needed |
| group_members | [GroupMembersSubset](#dialog.GroupMembersSubset) | repeated | Group &#43; subset of members to return. For loading members of big groups by chunks. |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated | groups needed |






<a name="dialog.RequestGetState"/>

### RequestGetState
Get main sequence state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| optimizations | [UpdateOptimization](#dialog.UpdateOptimization) | repeated |  |






<a name="dialog.RequestSubscribeFromGroupOnline"/>

### RequestSubscribeFromGroupOnline
Removing subscription for groups online


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.RequestSubscribeFromOnline"/>

### RequestSubscribeFromOnline
Removing subscription for users online


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.RequestSubscribeToGroupOnline"/>

### RequestSubscribeToGroupOnline
Subscribing for groups online


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |






<a name="dialog.RequestSubscribeToOnline"/>

### RequestSubscribeToOnline
Subscribing for users online


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.ResponseGetDialogsDifference"/>

### ResponseGetDialogsDifference
Dialogs &#43; peers and entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| dialogs | [Dialog](#dialog.Dialog) | repeated |  |
| group_peers | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.ResponseGetDifference"/>

### ResponseGetDifference
Updates happens after requested seq number &#43; related peers and entities


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  | seq of the last loaded update |
| state | [bytes](#bytes) |  |  |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| updates | [UpdateContainer](#dialog.UpdateContainer) | repeated |  |
| messages | [HistoryMessage](#dialog.HistoryMessage) | repeated |  |
| need_more | [bool](#bool) |  | false if all updates returned |
| users_refs | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| groups_refs | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| config | [Config](#dialog.Config) |  | user&#39;s config |
| config_hash | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |






<a name="dialog.ResponseGetReferencedEntitites"/>

### ResponseGetReferencedEntitites



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |
| messages | [HistoryMessage](#dialog.HistoryMessage) | repeated |  |






<a name="dialog.SeqUpdateBox"/>

### SeqUpdateBox
Container which contains UpdateSeqUpdate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| update | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  | serialized UpdateSeqUpdate |






<a name="dialog.UpdateCombinedUpdate"/>

### UpdateCombinedUpdate
Combined update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq_start | [int32](#int32) |  | start of related sequence intervals |
| seq_end | [int32](#int32) |  | end of related sequence intervals |
| state | [bytes](#bytes) |  |  |
| users | [User](#dialog.User) | repeated | related users |
| groups | [Group](#dialog.Group) | repeated | related groups |
| updates | [UpdateContainer](#dialog.UpdateContainer) | repeated |  |
| messages | [HistoryMessage](#dialog.HistoryMessage) | repeated | related messages |






<a name="dialog.UpdateContainer"/>

### UpdateContainer
Update container


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| update_header | [int32](#int32) |  |  |
| update | [bytes](#bytes) |  |  |






<a name="dialog.UpdateEmptyUpdate"/>

### UpdateEmptyUpdate
Empty update






<a name="dialog.UpdateFatSeqUpdate"/>

### UpdateFatSeqUpdate
Fat sequence update with additional data


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| update_header | [int32](#int32) |  |  |
| update | [bytes](#bytes) |  |  |
| users | [User](#dialog.User) | repeated | related users |
| groups | [Group](#dialog.Group) | repeated | related groups |






<a name="dialog.UpdateRawUpdate"/>

### UpdateRawUpdate
Custom Raw Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| bytes | [bytes](#bytes) |  |  |






<a name="dialog.UpdateSeqUpdate"/>

### UpdateSeqUpdate
Sequence update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| seq | [int32](#int32) |  |  |
| state | [bytes](#bytes) |  |  |
| update_header | [int32](#int32) |  |  |
| updateForceReloadState | [UpdateForceReloadState](#dialog.UpdateForceReloadState) |  |  |
| updateUserAvatarChanged | [UpdateUserAvatarChanged](#dialog.UpdateUserAvatarChanged) |  |  |
| updateUserNameChanged | [UpdateUserNameChanged](#dialog.UpdateUserNameChanged) |  |  |
| updateUserLocalNameChanged | [UpdateUserLocalNameChanged](#dialog.UpdateUserLocalNameChanged) |  |  |
| updateUserContactsChanged | [UpdateUserContactsChanged](#dialog.UpdateUserContactsChanged) |  |  |
| updateUserNickChanged | [UpdateUserNickChanged](#dialog.UpdateUserNickChanged) |  |  |
| updateUserAboutChanged | [UpdateUserAboutChanged](#dialog.UpdateUserAboutChanged) |  |  |
| updateUserPreferredLanguagesChanged | [UpdateUserPreferredLanguagesChanged](#dialog.UpdateUserPreferredLanguagesChanged) |  |  |
| updateUserTimeZoneChanged | [UpdateUserTimeZoneChanged](#dialog.UpdateUserTimeZoneChanged) |  |  |
| updateUserBotCommandsChanged | [UpdateUserBotCommandsChanged](#dialog.UpdateUserBotCommandsChanged) |  |  |
| updateUserExtChanged | [UpdateUserExtChanged](#dialog.UpdateUserExtChanged) |  |  |
| updateUserFullExtChanged | [UpdateUserFullExtChanged](#dialog.UpdateUserFullExtChanged) |  |  |
| updateUserSexChanged | [UpdateUserSexChanged](#dialog.UpdateUserSexChanged) |  |  |
| updateUserCustomProfileChanged | [UpdateUserCustomProfileChanged](#dialog.UpdateUserCustomProfileChanged) |  |  |
| updateUserStatusChanged | [UpdateUserStatusChanged](#dialog.UpdateUserStatusChanged) |  |  |
| updateContactRegistered | [UpdateContactRegistered](#dialog.UpdateContactRegistered) |  |  |
| updateContactsAdded | [UpdateContactsAdded](#dialog.UpdateContactsAdded) |  |  |
| updateContactsAddTaskSuspended | [UpdateContactsAddTaskSuspended](#dialog.UpdateContactsAddTaskSuspended) |  |  |
| updateContactsRemoved | [UpdateContactsRemoved](#dialog.UpdateContactsRemoved) |  |  |
| updateUserBlocked | [UpdateUserBlocked](#dialog.UpdateUserBlocked) |  |  |
| updateUserUnblocked | [UpdateUserUnblocked](#dialog.UpdateUserUnblocked) |  |  |
| updateInteractiveMediaEvent | [UpdateInteractiveMediaEvent](#dialog.UpdateInteractiveMediaEvent) |  |  |
| updateMessage | [UpdateMessage](#dialog.UpdateMessage) |  |  |
| updateMessageContentChanged | [UpdateMessageContentChanged](#dialog.UpdateMessageContentChanged) |  |  |
| updateMessageSent | [UpdateMessageSent](#dialog.UpdateMessageSent) |  |  |
| updateMessageReceived | [UpdateMessageReceived](#dialog.UpdateMessageReceived) |  |  |
| updateMessageRead | [UpdateMessageRead](#dialog.UpdateMessageRead) |  |  |
| updateMessageReadByMe | [UpdateMessageReadByMe](#dialog.UpdateMessageReadByMe) |  |  |
| updateMessageDelete | [UpdateMessageDelete](#dialog.UpdateMessageDelete) |  |  |
| updateChatClear | [UpdateChatClear](#dialog.UpdateChatClear) |  |  |
| updateChatDelete | [UpdateChatDelete](#dialog.UpdateChatDelete) |  |  |
| updateChatArchive | [UpdateChatArchive](#dialog.UpdateChatArchive) |  |  |
| updateChatGroupsChanged | [UpdateChatGroupsChanged](#dialog.UpdateChatGroupsChanged) |  |  |
| updateReactionsUpdate | [UpdateReactionsUpdate](#dialog.UpdateReactionsUpdate) |  |  |
| updateDialogFavouriteChanged | [UpdateDialogFavouriteChanged](#dialog.UpdateDialogFavouriteChanged) |  |  |
| updatePinnedMessagesChanged | [UpdatePinnedMessagesChanged](#dialog.UpdatePinnedMessagesChanged) |  |  |
| updateGroupTitleChanged | [UpdateGroupTitleChanged](#dialog.UpdateGroupTitleChanged) |  |  |
| updateGroupAvatarChanged | [UpdateGroupAvatarChanged](#dialog.UpdateGroupAvatarChanged) |  |  |
| updateGroupTopicChanged | [UpdateGroupTopicChanged](#dialog.UpdateGroupTopicChanged) |  |  |
| updateGroupAboutChanged | [UpdateGroupAboutChanged](#dialog.UpdateGroupAboutChanged) |  |  |
| updateGroupOwnerChanged | [UpdateGroupOwnerChanged](#dialog.UpdateGroupOwnerChanged) |  |  |
| updateGroupHistoryShared | [UpdateGroupHistoryShared](#dialog.UpdateGroupHistoryShared) |  |  |
| updateGroupCanSendMessagesChanged | [UpdateGroupCanSendMessagesChanged](#dialog.UpdateGroupCanSendMessagesChanged) |  |  |
| updateGroupCanViewMembersChanged | [UpdateGroupCanViewMembersChanged](#dialog.UpdateGroupCanViewMembersChanged) |  |  |
| updateGroupCanInviteMembersChanged | [UpdateGroupCanInviteMembersChanged](#dialog.UpdateGroupCanInviteMembersChanged) |  |  |
| updateGroupMemberChanged | [UpdateGroupMemberChanged](#dialog.UpdateGroupMemberChanged) |  |  |
| updateGroupMembersBecameAsync | [UpdateGroupMembersBecameAsync](#dialog.UpdateGroupMembersBecameAsync) |  |  |
| updateGroupMembersUpdated | [UpdateGroupMembersUpdated](#dialog.UpdateGroupMembersUpdated) |  |  |
| updateGroupMemberDiff | [UpdateGroupMemberDiff](#dialog.UpdateGroupMemberDiff) |  |  |
| updateGroupMembersCountChanged | [UpdateGroupMembersCountChanged](#dialog.UpdateGroupMembersCountChanged) |  |  |
| updateGroupMemberAdminChanged | [UpdateGroupMemberAdminChanged](#dialog.UpdateGroupMemberAdminChanged) |  |  |
| updateGroupMemberPermissionsChanged | [UpdateGroupMemberPermissionsChanged](#dialog.UpdateGroupMemberPermissionsChanged) |  |  |
| updateGroupInviteObsolete | [UpdateGroupInviteObsolete](#dialog.UpdateGroupInviteObsolete) |  |  |
| updateGroupUserInvitedObsolete | [UpdateGroupUserInvitedObsolete](#dialog.UpdateGroupUserInvitedObsolete) |  |  |
| updateGroupUserLeaveObsolete | [UpdateGroupUserLeaveObsolete](#dialog.UpdateGroupUserLeaveObsolete) |  |  |
| updateGroupUserKickObsolete | [UpdateGroupUserKickObsolete](#dialog.UpdateGroupUserKickObsolete) |  |  |
| updateGroupMembersUpdateObsolete | [UpdateGroupMembersUpdateObsolete](#dialog.UpdateGroupMembersUpdateObsolete) |  |  |
| updateGroupTitleChangedObsolete | [UpdateGroupTitleChangedObsolete](#dialog.UpdateGroupTitleChangedObsolete) |  |  |
| updateGroupTopicChangedObsolete | [UpdateGroupTopicChangedObsolete](#dialog.UpdateGroupTopicChangedObsolete) |  |  |
| updateGroupAboutChangedObsolete | [UpdateGroupAboutChangedObsolete](#dialog.UpdateGroupAboutChangedObsolete) |  |  |
| updateGroupAvatarChangedObsolete | [UpdateGroupAvatarChangedObsolete](#dialog.UpdateGroupAvatarChangedObsolete) |  |  |
| updateGroupShortnameChanged | [UpdateGroupShortnameChanged](#dialog.UpdateGroupShortnameChanged) |  |  |
| updateStickerCollectionsChanged | [UpdateStickerCollectionsChanged](#dialog.UpdateStickerCollectionsChanged) |  |  |
| updateStickerPackRemoved | [UpdateStickerPackRemoved](#dialog.UpdateStickerPackRemoved) |  |  |
| updateStickerPackAdded | [UpdateStickerPackAdded](#dialog.UpdateStickerPackAdded) |  |  |
| updatePauseNotifications | [UpdatePauseNotifications](#dialog.UpdatePauseNotifications) |  |  |
| updateRestoreNotifications | [UpdateRestoreNotifications](#dialog.UpdateRestoreNotifications) |  |  |
| updateTyping | [UpdateTyping](#dialog.UpdateTyping) |  |  |
| updateTypingStop | [UpdateTypingStop](#dialog.UpdateTypingStop) |  |  |
| updateUserOnline | [UpdateUserOnline](#dialog.UpdateUserOnline) |  |  |
| updateUserOffline | [UpdateUserOffline](#dialog.UpdateUserOffline) |  |  |
| updateUserLastSeen | [UpdateUserLastSeen](#dialog.UpdateUserLastSeen) |  |  |
| updateGroupOnline | [UpdateGroupOnline](#dialog.UpdateGroupOnline) |  |  |
| updateEventBusDeviceConnected | [UpdateEventBusDeviceConnected](#dialog.UpdateEventBusDeviceConnected) |  |  |
| updateEventBusDeviceDisconnected | [UpdateEventBusDeviceDisconnected](#dialog.UpdateEventBusDeviceDisconnected) |  |  |
| updateEventBusMessage | [UpdateEventBusMessage](#dialog.UpdateEventBusMessage) |  |  |
| updateEventBusDisposed | [UpdateEventBusDisposed](#dialog.UpdateEventBusDisposed) |  |  |
| updateIncomingCallDeprecated | [UpdateIncomingCallDeprecated](#dialog.UpdateIncomingCallDeprecated) |  |  |
| updateIncomingCall | [UpdateIncomingCall](#dialog.UpdateIncomingCall) |  |  |
| updateCallHandled | [UpdateCallHandled](#dialog.UpdateCallHandled) |  |  |
| updateCallDisposed | [UpdateCallDisposed](#dialog.UpdateCallDisposed) |  |  |
| updateParameterChanged | [UpdateParameterChanged](#dialog.UpdateParameterChanged) |  |  |
| updateRawUpdate | [UpdateRawUpdate](#dialog.UpdateRawUpdate) |  |  |
| updateEmptyUpdate | [UpdateEmptyUpdate](#dialog.UpdateEmptyUpdate) |  |  |
| updateCountersChanged | [UpdateCountersChanged](#dialog.UpdateCountersChanged) |  |  |
| updateConfig | [UpdateConfig](#dialog.UpdateConfig) |  |  |






<a name="dialog.UpdateSeqUpdateTooLong"/>

### UpdateSeqUpdateTooLong
Notification about requiring performing manual GetDifference






<a name="dialog.UpdateWeakFatUpdate"/>

### UpdateWeakFatUpdate
Fat Weak Update


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [int64](#int64) |  |  |
| update_header | [int32](#int32) |  |  |
| update | [bytes](#bytes) |  |  |
| users | [User](#dialog.User) | repeated |  |
| groups | [Group](#dialog.Group) | repeated |  |






<a name="dialog.UpdateWeakUpdate"/>

### UpdateWeakUpdate
Out of sequence update (for typing and online statuses)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| date | [int64](#int64) |  |  |
| update_header | [int32](#int32) |  |  |
| update | [bytes](#bytes) |  |  |





 

 

 


<a name="dialog.SequenceAndUpdates"/>

### SequenceAndUpdates


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetState | [RequestGetState](#dialog.RequestGetState) | [ResponseSeq](#dialog.RequestGetState) | Get last seq number |
| GetDifference | [RequestGetDifference](#dialog.RequestGetDifference) | [ResponseGetDifference](#dialog.RequestGetDifference) | Get all update that happens after given seq number |
| GetDialogsDifference | [RequestGetDialogsDifference](#dialog.RequestGetDialogsDifference) | [ResponseGetDialogsDifference](#dialog.RequestGetDialogsDifference) | Load all dialogs that changed after given date |
| GetReferencedEntitites | [RequestGetReferencedEntitites](#dialog.RequestGetReferencedEntitites) | [ResponseGetReferencedEntitites](#dialog.RequestGetReferencedEntitites) | Load some required entities |
| SubscribeToOnline | [RequestSubscribeToOnline](#dialog.RequestSubscribeToOnline) | [ResponseVoid](#dialog.RequestSubscribeToOnline) |  |
| SubscribeFromOnline | [RequestSubscribeFromOnline](#dialog.RequestSubscribeFromOnline) | [ResponseVoid](#dialog.RequestSubscribeFromOnline) |  |
| SubscribeToGroupOnline | [RequestSubscribeToGroupOnline](#dialog.RequestSubscribeToGroupOnline) | [ResponseVoid](#dialog.RequestSubscribeToGroupOnline) |  |
| SubscribeFromGroupOnline | [RequestSubscribeFromGroupOnline](#dialog.RequestSubscribeFromGroupOnline) | [ResponseVoid](#dialog.RequestSubscribeFromGroupOnline) |  |
| SeqUpdates | [.google.protobuf.Empty](#google.protobuf.Empty) | [SeqUpdateBox](#google.protobuf.Empty) | Get stream of the user&#39;s updates |

 



<a name="stickers.proto"/>
<p align="right"><a href="#top">Top</a></p>

## stickers.proto



<a name="dialog.RequestAddStickerCollection"/>

### RequestAddStickerCollection
Adding sticker collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  |  |






<a name="dialog.RequestAddStickerPackReference"/>

### RequestAddStickerPackReference
Add a reference to other user&#39;s sticker pack


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_sticker_pack | [int32](#int32) |  |  |






<a name="dialog.RequestLoadAcesssibleStickers"/>

### RequestLoadAcesssibleStickers
Load accessible stickers






<a name="dialog.RequestLoadOwnStickers"/>

### RequestLoadOwnStickers
Loading own stickers






<a name="dialog.RequestLoadStickerCollection"/>

### RequestLoadStickerCollection
Loading stickers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="dialog.RequestRemoveStickerCollection"/>

### RequestRemoveStickerCollection
Removing sticker collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  |  |






<a name="dialog.RequestRemoveStickerPackReference"/>

### RequestRemoveStickerPackReference
Remove a reference to an other user&#39;s sticker pack


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| source_sticker_pack | [int32](#int32) |  |  |






<a name="dialog.ResponseLoadAcesssibleStickers"/>

### ResponseLoadAcesssibleStickers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| accessible_stickers | [StickerCollection](#dialog.StickerCollection) | repeated |  |






<a name="dialog.ResponseLoadOwnStickers"/>

### ResponseLoadOwnStickers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| own_stickers | [StickerCollection](#dialog.StickerCollection) | repeated |  |






<a name="dialog.ResponseLoadStickerCollection"/>

### ResponseLoadStickerCollection



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collection | [StickerCollection](#dialog.StickerCollection) |  |  |






<a name="dialog.ResponseStickersResponse"/>

### ResponseStickersResponse
Stickers response


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| collections | [StickerCollection](#dialog.StickerCollection) | repeated |  |
| seq | [int32](#int32) |  | deprecated |
| state | [bytes](#bytes) |  |  |






<a name="dialog.StickerCollection"/>

### StickerCollection
Sticker collection


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Unique id of a collection |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Sticker pack title |
| stickers | [StickerDescriptor](#dialog.StickerDescriptor) | repeated |  |
| owned_by_me | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | does this pack belongs to current user |






<a name="dialog.StickerDescriptor"/>

### StickerDescriptor
Descriptor of a Sticker


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | Sticker unique id |
| emoji | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Emoji code for sticker |
| image_128 | [ImageLocation](#dialog.ImageLocation) |  | Image of sticker 128x128 in WebP format |
| image_512 | [ImageLocation](#dialog.ImageLocation) |  | Image of sticker 512x512 in WebP format |
| image_256 | [ImageLocation](#dialog.ImageLocation) |  | Image of sticker 256x256 in WebP format |






<a name="dialog.UpdateStickerCollectionsChanged"/>

### UpdateStickerCollectionsChanged
Sticker collection changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| updated_collections | [StickerCollection](#dialog.StickerCollection) | repeated |  |






<a name="dialog.UpdateStickerPackAdded"/>

### UpdateStickerPackAdded
Sticker pack was added


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack | [StickerCollection](#dialog.StickerCollection) |  |  |






<a name="dialog.UpdateStickerPackRemoved"/>

### UpdateStickerPackRemoved
Sticker pack removed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| pack_id | [int32](#int32) |  |  |





 

 

 


<a name="dialog.Stickers"/>

### Stickers


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| LoadOwnStickers | [RequestLoadOwnStickers](#dialog.RequestLoadOwnStickers) | [ResponseLoadOwnStickers](#dialog.RequestLoadOwnStickers) |  |
| LoadAcesssibleStickers | [RequestLoadAcesssibleStickers](#dialog.RequestLoadAcesssibleStickers) | [ResponseLoadAcesssibleStickers](#dialog.RequestLoadAcesssibleStickers) |  |
| AddStickerPackReference | [RequestAddStickerPackReference](#dialog.RequestAddStickerPackReference) | [ResponseSeq](#dialog.RequestAddStickerPackReference) |  |
| RemoveStickerPackReference | [RequestRemoveStickerPackReference](#dialog.RequestRemoveStickerPackReference) | [ResponseSeq](#dialog.RequestRemoveStickerPackReference) |  |
| AddStickerCollection | [RequestAddStickerCollection](#dialog.RequestAddStickerCollection) | [ResponseSeq](#dialog.RequestAddStickerCollection) |  |
| RemoveStickerCollection | [RequestRemoveStickerCollection](#dialog.RequestRemoveStickerCollection) | [ResponseSeq](#dialog.RequestRemoveStickerCollection) |  |
| LoadStickerCollection | [RequestLoadStickerCollection](#dialog.RequestLoadStickerCollection) | [ResponseLoadStickerCollection](#dialog.RequestLoadStickerCollection) |  |

 



<a name="typing_and_online.proto"/>
<p align="right"><a href="#top">Top</a></p>

## typing_and_online.proto



<a name="dialog.RequestPauseNotifications"/>

### RequestPauseNotifications
Pause notifications


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timeout | [int32](#int32) |  |  |






<a name="dialog.RequestRestoreNotifications"/>

### RequestRestoreNotifications
Restoring notifications






<a name="dialog.RequestSetOnline"/>

### RequestSetOnline
Sending online state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_online | [bool](#bool) |  |  |
| timeout | [int64](#int64) |  | offline after timeout |
| device_type | [DeviceType](#dialog.DeviceType) |  |  |
| device_category | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestStopTyping"/>

### RequestStopTyping
Stop typing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| typing_type | [TypingType](#dialog.TypingType) |  |  |






<a name="dialog.RequestTyping"/>

### RequestTyping
Sending typing notification


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| typing_type | [TypingType](#dialog.TypingType) |  |  |






<a name="dialog.UpdateGroupOnline"/>

### UpdateGroupOnline
Update about group online change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int32](#int32) |  |  |
| count | [int32](#int32) |  | amount of online users |






<a name="dialog.UpdatePauseNotifications"/>

### UpdatePauseNotifications
Update about pausing notifications


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| timeout | [int32](#int32) |  |  |






<a name="dialog.UpdateRestoreNotifications"/>

### UpdateRestoreNotifications
Update about restoring notifications






<a name="dialog.UpdateTyping"/>

### UpdateTyping
Update about user&#39;s typing


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| uid | [int32](#int32) |  |  |
| typing_type | [TypingType](#dialog.TypingType) |  |  |






<a name="dialog.UpdateTypingStop"/>

### UpdateTypingStop
Update about user&#39;s typing stop


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| uid | [int32](#int32) |  |  |
| typing_type | [TypingType](#dialog.TypingType) |  |  |






<a name="dialog.UpdateUserLastSeen"/>

### UpdateUserLastSeen
Update about user&#39;s last seen state


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| date | [int64](#int64) |  |  |
| device_type | [DeviceType](#dialog.DeviceType) |  |  |
| device_category | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateUserOffline"/>

### UpdateUserOffline
Update about user became offline


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| device_type | [DeviceType](#dialog.DeviceType) |  |  |
| device_category | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateUserOnline"/>

### UpdateUserOnline
Update about user became online


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| device_type | [DeviceType](#dialog.DeviceType) |  |  |
| device_category | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |





 


<a name="dialog.DeviceType"/>

### DeviceType


| Name | Number | Description |
| ---- | ------ | ----------- |
| DEVICETYPE_UNKNOWN | 0 |  |
| DEVICETYPE_GENERIC | 1 |  |
| DEVICETYPE_PC | 2 |  |
| DEVICETYPE_MOBILE | 3 |  |
| DEVICETYPE_TABLET | 4 |  |
| DEVICETYPE_WATCH | 5 |  |
| DEVICETYPE_MIRROR | 6 |  |
| DEVICETYPE_CAR | 7 |  |
| DEVICETYPE_TABLE | 8 |  |



<a name="dialog.TypingType"/>

### TypingType


| Name | Number | Description |
| ---- | ------ | ----------- |
| TYPINGTYPE_UNKNOWN | 0 |  |
| TYPINGTYPE_TEXT | 1 |  |


 

 


<a name="dialog.TypingAndOnline"/>

### TypingAndOnline


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Typing | [RequestTyping](#dialog.RequestTyping) | [ResponseVoid](#dialog.RequestTyping) |  |
| StopTyping | [RequestStopTyping](#dialog.RequestStopTyping) | [ResponseVoid](#dialog.RequestStopTyping) |  |
| SetOnline | [RequestSetOnline](#dialog.RequestSetOnline) | [ResponseVoid](#dialog.RequestSetOnline) |  |
| PauseNotifications | [RequestPauseNotifications](#dialog.RequestPauseNotifications) | [ResponseVoid](#dialog.RequestPauseNotifications) |  |
| RestoreNotifications | [RequestRestoreNotifications](#dialog.RequestRestoreNotifications) | [ResponseVoid](#dialog.RequestRestoreNotifications) |  |

 



<a name="users.proto"/>
<p align="right"><a href="#top">Top</a></p>

## users.proto



<a name="dialog.BotCommand"/>

### BotCommand
Available bot commands


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| slash_command | [string](#string) |  | Slash command name (wihtout slash) |
| description | [string](#string) |  | Slash command description |
| loc_key | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Optional Localization Key for i18n |






<a name="dialog.ContactRecord"/>

### ContactRecord
Contact information record


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [ContactType](#dialog.ContactType) |  |  |
| type_spec | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Value for specification type of contact, for example &#39;mobile/standalone/office&#39; for phones or &#39;vk/fb/telegram&#39; for extenrnal networks. |
| string_value | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| long_value | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |
| title | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| subtitle | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.FullUser"/>

### FullUser
Full User representation - deprecated


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | User&#39;s id |
| contact_info | [ContactRecord](#dialog.ContactRecord) | repeated |  |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| preferred_languages | [string](#string) | repeated |  |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Time Zone of user in TZ format |
| bot_commands | [BotCommand](#dialog.BotCommand) | repeated |  |
| is_blocked | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| custom_profile | [string](#string) |  | custom user profile info in JSON format |
| integration_token | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| status | [UserStatus](#dialog.UserStatus) |  |  |






<a name="dialog.RequestEditUserLocalName"/>

### RequestEditUserLocalName
Renaming of user&#39;s visible name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| access_hash | [int64](#int64) |  |  |
| name | [string](#string) |  |  |






<a name="dialog.RequestLoadFullUsers"/>

### RequestLoadFullUsers
Loading Full User information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_peers | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |






<a name="dialog.ResponseLoadFullUsers"/>

### ResponseLoadFullUsers



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| full_users | [FullUser](#dialog.FullUser) | repeated |  |






<a name="dialog.UpdateUserAboutChanged"/>

### UpdateUserAboutChanged
Update about user&#39;s about changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| about | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateUserAvatarChanged"/>

### UpdateUserAvatarChanged
Update about avatar changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| avatar | [Avatar](#dialog.Avatar) |  |  |






<a name="dialog.UpdateUserBotCommandsChanged"/>

### UpdateUserBotCommandsChanged
Update about bot commands changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| commands | [BotCommand](#dialog.BotCommand) | repeated |  |






<a name="dialog.UpdateUserContactsChanged"/>

### UpdateUserContactsChanged
Update about contact information change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| contact_records | [ContactRecord](#dialog.ContactRecord) | repeated |  |






<a name="dialog.UpdateUserCustomProfileChanged"/>

### UpdateUserCustomProfileChanged
Update about user custom profile changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| custom_profile | [string](#string) |  | custom user profile info in JSON format |






<a name="dialog.UpdateUserExtChanged"/>

### UpdateUserExtChanged



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| ext | [RecursiveMapValue](#dialog.RecursiveMapValue) |  |  |






<a name="dialog.UpdateUserFullExtChanged"/>

### UpdateUserFullExtChanged



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| ext | [RecursiveMapValue](#dialog.RecursiveMapValue) |  |  |






<a name="dialog.UpdateUserLocalNameChanged"/>

### UpdateUserLocalNameChanged
Update about local name changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| local_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateUserNameChanged"/>

### UpdateUserNameChanged
Update about name changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| name | [string](#string) |  |  |






<a name="dialog.UpdateUserNickChanged"/>

### UpdateUserNickChanged
Update about nick changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| nickname | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.UpdateUserPreferredLanguagesChanged"/>

### UpdateUserPreferredLanguagesChanged
Update about user&#39;s preferred languages


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| preferred_languages | [string](#string) | repeated |  |






<a name="dialog.UpdateUserSexChanged"/>

### UpdateUserSexChanged
Update about user sex changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| sex | [Sex](#dialog.Sex) |  |  |






<a name="dialog.UpdateUserStatusChanged"/>

### UpdateUserStatusChanged
Update about user status change


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| status | [UserStatus](#dialog.UserStatus) |  |  |






<a name="dialog.UpdateUserTimeZoneChanged"/>

### UpdateUserTimeZoneChanged
User TimeZone changed


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| uid | [int32](#int32) |  |  |
| time_zone | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.User"/>

### User
Main user object


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | User&#39;s id |
| access_hash | [int64](#int64) |  |  |
| name | [string](#string) |  |  |
| local_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| nick | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |
| sex | [Sex](#dialog.Sex) |  |  |
| avatar | [Avatar](#dialog.Avatar) |  |  |
| is_bot | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |






<a name="dialog.UserStatus"/>

### UserStatus
User&#39;s status


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [UserStatusType](#dialog.UserStatusType) |  |  |
| text | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Text supplied by user |





 


<a name="dialog.ContactType"/>

### ContactType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CONTACTTYPE_UNKNOWN | 0 |  |
| CONTACTTYPE_PHONE | 1 |  |
| CONTACTTYPE_EMAIL | 2 |  |
| CONTACTTYPE_WEB | 3 |  |
| CONTACTTYPE_SOCIAL | 4 |  |



<a name="dialog.Sex"/>

### Sex


| Name | Number | Description |
| ---- | ------ | ----------- |
| SEX_UNKNOWN | 0 |  |
| SEX_MALE | 2 |  |
| SEX_FEMALE | 3 |  |



<a name="dialog.UserStatusType"/>

### UserStatusType


| Name | Number | Description |
| ---- | ------ | ----------- |
| USERSTATUSTYPE_UNKNOWN | 0 |  |
| USERSTATUSTYPE_UNSET | 1 |  |
| USERSTATUSTYPE_AWAY | 2 |  |
| USERSTATUSTYPE_DONOTDISTURB | 3 |  |
| USERSTATUSTYPE_INVISIBLE | 4 |  |


 

 


<a name="dialog.Users"/>

### Users


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| EditUserLocalName | [RequestEditUserLocalName](#dialog.RequestEditUserLocalName) | [ResponseSeq](#dialog.RequestEditUserLocalName) |  |
| LoadFullUsers | [RequestLoadFullUsers](#dialog.RequestLoadFullUsers) | [ResponseLoadFullUsers](#dialog.RequestLoadFullUsers) | Deprecated |

 



<a name="web_rtc.proto"/>
<p align="right"><a href="#top">Top</a></p>

## web_rtc.proto



<a name="dialog.ActiveCall"/>

### ActiveCall
Active Calls. Used in broadcasting states of current calls.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| peer | [Peer](#dialog.Peer) |  | Call&#39;s peer |
| call_members | [CallMember](#dialog.CallMember) | repeated | Call Members |






<a name="dialog.AdvertiseMaster"/>

### AdvertiseMaster
Sent by master


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| server | [ICEServer](#dialog.ICEServer) | repeated | a list of ICE servers |
| call_api_version | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  | a verison of the call signaling |
| call_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.AdvertisePeer"/>

### AdvertisePeer
Peer advertisement


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device_id | [int64](#int64) |  | device id of the peer |
| peer_settings | [PeerSettings](#dialog.PeerSettings) |  | settings for the peer |
| ice_servers | [ICEServer](#dialog.ICEServer) | repeated | ice servers for the peer |






<a name="dialog.AdvertiseSelf"/>

### AdvertiseSelf
Advertizing self to a master mode


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer_settings | [PeerSettings](#dialog.PeerSettings) |  | Optional peer Settings |






<a name="dialog.Answer"/>

### Answer
Answer signal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  | Session Id of answer |
| sdp | [string](#string) |  | Answer SDP |






<a name="dialog.CallLogEntry"/>

### CallLogEntry
This struct represents a call from the history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | call id |
| call_date | [int64](#int64) |  |  |
| initiator | [OutPeer](#dialog.OutPeer) |  |  |
| recipient | [OutPeer](#dialog.OutPeer) |  |  |
| duration | [google.protobuf.Int64Value](#google.protobuf.Int64Value) |  |  |
| answered | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | was the call answered by anybody |
| finished | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |
| display_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | optional display name for this call |
| rejected | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  |  |






<a name="dialog.CallMember"/>

### CallMember
Call Member


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_id | [int32](#int32) |  | Member User Id |
| state | [CallMemberStateHolder](#dialog.CallMemberStateHolder) |  | State of member |






<a name="dialog.CallMemberStateHolder"/>

### CallMemberStateHolder
Call Member state holder


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| state | [CallMemberState](#dialog.CallMemberState) |  | State Value |
| fallback_is_ringing | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Fallback flag for future compatibility of state |
| fallback_is_connected | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Fallback flag for future compatibility of state |
| fallback_is_connecting | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Fallback flag for future compatibility of state |
| fallback_is_ringing_reached | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Fallback flag for future compatibility of state |
| fallback_is_ended | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Fallback flag for future compatibility of state |






<a name="dialog.CallNameChanged"/>

### CallNameChanged
Used to notify participants about call name changes


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.CallStats"/>

### CallStats
Used to send call statistics to the server


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [CallStatsType](#dialog.CallStatsType) |  | the type of the event |
| value | [string](#string) |  | the value of the event |






<a name="dialog.Candidate"/>

### Candidate
Candidate signal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  | Session Id of candidate |
| index | [int32](#int32) |  | Index of candidate |
| id | [string](#string) |  | Id of candidate |
| sdp | [string](#string) |  | SDP of candidate |






<a name="dialog.CloseSession"/>

### CloseSession
Close this session and be ready to


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Device Id |
| session_id | [int64](#int64) |  | Session Id for renegotiation |






<a name="dialog.DTMF"/>

### DTMF
Emulates DTMF key press


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| code | [DTMFCode](#dialog.DTMFCode) |  |  |
| duration | [int64](#int64) |  | the duration (in milliseconds) |






<a name="dialog.EnableConnection"/>

### EnableConnection
Notification about enabling connection to peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Device Id |
| session_id | [int64](#int64) |  | Session Id |






<a name="dialog.GotICECandidate"/>

### GotICECandidate
WebRTC IceCandidate add signal.
Used to let the other party know about new ICE candidate, usually from onIceCandidate callback of WebRTC


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  |  |
| candidate | [ICECandidate](#dialog.ICECandidate) |  |  |






<a name="dialog.ICECandidate"/>

### ICECandidate
Structure representing an ICE candidate


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| sdp_m_line_index | [int32](#int32) |  |  |
| sdp_mid | [string](#string) |  |  |
| sdp | [string](#string) |  |  |






<a name="dialog.ICEServer"/>

### ICEServer
ICE Server description


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  | Url to server |
| username | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Optional username |
| credential | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Optional credential |






<a name="dialog.NeedDisconnect"/>

### NeedDisconnect
Notification about requirement about required disconnection from peer


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Device Id |






<a name="dialog.NeedOffer"/>

### NeedOffer
Notification from master that offer is required


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Destination Device Id |
| session_id | [int64](#int64) |  | Session Id |
| peer_settings | [PeerSettings](#dialog.PeerSettings) |  | Deprecated peer settings |






<a name="dialog.NegotinationSuccessful"/>

### NegotinationSuccessful
Notification about on negotiation is successful


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Device Id |
| session_id | [int64](#int64) |  | Session Id |






<a name="dialog.Offer"/>

### Offer
Offer signal


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  | Session Id of offer |
| sdp | [string](#string) |  | Offer SDP |
| peer_settings | [PeerSettings](#dialog.PeerSettings) |  | Deprecated Own Peer settings |






<a name="dialog.OnRenegotiationNeeded"/>

### OnRenegotiationNeeded
Need renegotiate session. For example when streams are changed.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| device | [int64](#int64) |  | Device Id |
| session_id | [int64](#int64) |  | Session Id |






<a name="dialog.PeerSettings"/>

### PeerSettings
Peer Settings


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| can_pre_connect | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | If peer can preconnect before answer |
| sends_offer | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | The peer already has an offer. The server should request offer from this peer. |
| wait_for_candidates | [google.protobuf.BoolValue](#google.protobuf.BoolValue) |  | Peer prefers to receive an offer/answer with ICE candidates collected |
| client_info | [ClientInfo](#dialog.ClientInfo) |  | Description of this peer. Version, browser, OS, type, etc. |






<a name="dialog.RemovedICECandidates"/>

### RemovedICECandidates
WebRTC IceCandidate remove signal
Used to let the other party know about ICE candidate list change, usually from onIceCandidatesRemoved callback


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| session_id | [int64](#int64) |  |  |
| candidate | [ICECandidate](#dialog.ICECandidate) | repeated |  |






<a name="dialog.RequestChangeCallDisplayName"/>

### RequestChangeCallDisplayName
Changes the call display name


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| display_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.RequestDeleteCall"/>

### RequestDeleteCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |






<a name="dialog.RequestDoCall"/>

### RequestDoCall
Do Call. Right after a call client need to start sending CallInProgress


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [OutPeer](#dialog.OutPeer) |  |  |
| timeout | [int64](#int64) |  |  |






<a name="dialog.RequestGetCallInfo"/>

### RequestGetCallInfo
Getting Call Information


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |






<a name="dialog.RequestJoinCall"/>

### RequestJoinCall
Joining Call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |






<a name="dialog.RequestLoadCalls"/>

### RequestLoadCalls
Call this function to load call history


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| next_offset | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  | Cursor |
| limit | [int32](#int32) |  |  |






<a name="dialog.RequestRejectCall"/>

### RequestRejectCall
Rejecting Call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| reason | [RejectCallReason](#dialog.RejectCallReason) |  |  |






<a name="dialog.ResponseDoCall"/>

### ResponseDoCall



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| event_bus_id | [string](#string) |  |  |
| device_id | [int64](#int64) |  |  |






<a name="dialog.ResponseGetCallInfo"/>

### ResponseGetCallInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| peer | [Peer](#dialog.Peer) |  |  |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| event_bus_id | [string](#string) |  |  |
| display_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  |  |






<a name="dialog.ResponseLoadCalls"/>

### ResponseLoadCalls
Calls &#43; related peers


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| calls | [CallLogEntry](#dialog.CallLogEntry) | repeated |  |
| users | [UserOutPeer](#dialog.UserOutPeer) | repeated |  |
| groups | [GroupOutPeer](#dialog.GroupOutPeer) | repeated |  |
| next_offset | [google.protobuf.BytesValue](#google.protobuf.BytesValue) |  |  |






<a name="dialog.UpdateCallDisposed"/>

### UpdateCallDisposed
Update about the call ending (all participants have left)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| attempt_index | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| reason | [CallDisposedReason](#dialog.CallDisposedReason) |  |  |






<a name="dialog.UpdateCallHandled"/>

### UpdateCallHandled
Update about incoming call handled


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| attempt_index | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |






<a name="dialog.UpdateIncomingCall"/>

### UpdateIncomingCall
A new update about incoming call


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| event_bus_id | [string](#string) |  |  |
| peer | [Peer](#dialog.Peer) |  |  |
| display_name | [google.protobuf.StringValue](#google.protobuf.StringValue) |  | Call visual name |
| attempt_index | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |
| out_peer | [OutPeer](#dialog.OutPeer) |  |  |
| video | [bool](#bool) |  |  |






<a name="dialog.UpdateIncomingCallDeprecated"/>

### UpdateIncomingCallDeprecated
Update about incoming call (Sent every 10 seconds)


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| call_id | [int64](#int64) |  |  |
| attempt_index | [google.protobuf.Int32Value](#google.protobuf.Int32Value) |  |  |






<a name="dialog.WebRTCSignaling"/>

### WebRTCSignaling



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| advertiseSelf | [AdvertiseSelf](#dialog.AdvertiseSelf) |  |  |
| advertiseMaster | [AdvertiseMaster](#dialog.AdvertiseMaster) |  |  |
| advertisePeer | [AdvertisePeer](#dialog.AdvertisePeer) |  |  |
| candidate | [Candidate](#dialog.Candidate) |  |  |
| gotICECandidate | [GotICECandidate](#dialog.GotICECandidate) |  |  |
| removedICECandidates | [RemovedICECandidates](#dialog.RemovedICECandidates) |  |  |
| offer | [Offer](#dialog.Offer) |  |  |
| answer | [Answer](#dialog.Answer) |  |  |
| needOffer | [NeedOffer](#dialog.NeedOffer) |  |  |
| negotinationSuccessful | [NegotinationSuccessful](#dialog.NegotinationSuccessful) |  |  |
| enableConnection | [EnableConnection](#dialog.EnableConnection) |  |  |
| onRenegotiationNeeded | [OnRenegotiationNeeded](#dialog.OnRenegotiationNeeded) |  |  |
| closeSession | [CloseSession](#dialog.CloseSession) |  |  |
| needDisconnect | [NeedDisconnect](#dialog.NeedDisconnect) |  |  |
| dTMF | [DTMF](#dialog.DTMF) |  |  |
| callStats | [CallStats](#dialog.CallStats) |  |  |
| callNameChanged | [CallNameChanged](#dialog.CallNameChanged) |  |  |





 


<a name="dialog.CallDisposedReason"/>

### CallDisposedReason


| Name | Number | Description |
| ---- | ------ | ----------- |
| CALLDISPOSEDREASON_UNKNOWN | 0 |  |
| CALLDISPOSEDREASON_REJECTED | 1 |  |
| CALLDISPOSEDREASON_BUSY | 2 |  |
| CALLDISPOSEDREASON_ENDED | 3 |  |
| CALLDISPOSEDREASON_ANSWER_TIMEOUT | 4 |  |



<a name="dialog.CallMemberState"/>

### CallMemberState


| Name | Number | Description |
| ---- | ------ | ----------- |
| CALLMEMBERSTATE_UNKNOWN | 0 |  |
| CALLMEMBERSTATE_RINGING | 1 |  |
| CALLMEMBERSTATE_RINGING_REACHED | 4 |  |
| CALLMEMBERSTATE_CONNECTING | 2 |  |
| CALLMEMBERSTATE_CONNECTED | 3 |  |
| CALLMEMBERSTATE_ENDED | 5 |  |



<a name="dialog.CallStatsType"/>

### CallStatsType


| Name | Number | Description |
| ---- | ------ | ----------- |
| CALLSTATSTYPE_UNKNOWN | 0 |  |
| CALLSTATSTYPE_TURN_SERVER_ROUND_TRIP_TIME | 1 |  |
| CALLSTATSTYPE_ICE_CONNECTION_FAILED | 2 |  |
| CALLSTATSTYPE_ICE_CONNECTION_CONNECTED | 3 |  |
| CALLSTATSTYPE_ICE_CONNECTION_DID_NOT_CONNECT | 4 |  |
| CALLSTATSTYPE_ICE_CONNECTION_CONNECTED_TIME | 5 |  |
| CALLSTATSTYPE_PEER_CONNECTION_CREATE_ANSWER_FAILED | 6 |  |
| CALLSTATSTYPE_PEER_CONNECTION_CREATE_OFFER_FAILED | 7 |  |
| CALLSTATSTYPE_PEER_CONNECTION_SET_REMOTE_DESCRIPTION_FAILED | 8 |  |
| CALLSTATSTYPE_PEER_CONNECTION_SET_LOCAL_DESCRIPTION_FAILED | 9 |  |



<a name="dialog.DTMFCode"/>

### DTMFCode


| Name | Number | Description |
| ---- | ------ | ----------- |
| DTMFCODE_UNKNOWN | 0 |  |
| DTMFCODE_ZERO | 1 |  |
| DTMFCODE_ONE | 2 |  |
| DTMFCODE_TWO | 3 |  |
| DTMFCODE_THREE | 4 |  |
| DTMFCODE_FOUR | 5 |  |
| DTMFCODE_FIVE | 6 |  |
| DTMFCODE_SIX | 7 |  |
| DTMFCODE_SEVEN | 8 |  |
| DTMFCODE_EIGHT | 9 |  |
| DTMFCODE_NINE | 10 |  |
| DTMFCODE_ASTERISK | 11 |  |
| DTMFCODE_POUND | 12 |  |
| DTMFCODE_A | 13 |  |
| DTMFCODE_B | 14 |  |
| DTMFCODE_C | 15 |  |
| DTMFCODE_D | 16 |  |



<a name="dialog.RejectCallReason"/>

### RejectCallReason


| Name | Number | Description |
| ---- | ------ | ----------- |
| REJECTCALLREASON_UNKNOWN | 0 |  |
| REJECTCALLREASON_DECLINE | 1 |  |
| REJECTCALLREASON_BUSY | 2 |  |


 

 


<a name="dialog.WebRTC"/>

### WebRTC


| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetCallInfo | [RequestGetCallInfo](#dialog.RequestGetCallInfo) | [ResponseGetCallInfo](#dialog.RequestGetCallInfo) |  |
| LoadCalls | [RequestLoadCalls](#dialog.RequestLoadCalls) | [ResponseLoadCalls](#dialog.RequestLoadCalls) |  |
| DoCall | [RequestDoCall](#dialog.RequestDoCall) | [ResponseDoCall](#dialog.RequestDoCall) |  |
| JoinCall | [RequestJoinCall](#dialog.RequestJoinCall) | [ResponseVoid](#dialog.RequestJoinCall) |  |
| RejectCall | [RequestRejectCall](#dialog.RequestRejectCall) | [ResponseVoid](#dialog.RequestRejectCall) |  |
| ChangeCallDisplayName | [RequestChangeCallDisplayName](#dialog.RequestChangeCallDisplayName) | [ResponseVoid](#dialog.RequestChangeCallDisplayName) |  |
| DeleteCall | [RequestDeleteCall](#dialog.RequestDeleteCall) | [ResponseVoid](#dialog.RequestDeleteCall) |  |

 



## Scalar Value Types

| .proto Type | Notes | C++ Type | Java Type | Python Type |
| ----------- | ----- | -------- | --------- | ----------- |
| <a name="double" /> double |  | double | double | float |
| <a name="float" /> float |  | float | float | float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long |
| <a name="bool" /> bool |  | bool | boolean | boolean |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str |

