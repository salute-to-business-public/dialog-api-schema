package dialog

import grpc "google.golang.org/grpc"
import metadata "google.golang.org/grpc/metadata"
import context "golang.org/x/net/context"
import . "github.com/golang/protobuf/ptypes/wrappers"
import . "github.com/golang/protobuf/ptypes/empty"
func (d *DialogSDK) initializeContext() error {
    ctx, cancel := context.WithCancel(context.Background())
    res, err := d.registration.RegisterDevice(ctx, &RequestRegisterDevice{})
    if err != nil {
        return err
    }
    d.internalContext = metadata.AppendToOutgoingContext(ctx, "x-auth-ticket", res.Token)
    d.cancel = cacnel
    return nil
}
func (d *DialogSDK) Close() {
    d.cancel()
    d.conn.Close()
}
type DialogSDK struct {
	internalContext context.Context
	cancel context.CancelFunc
	conn *grpc.ClientConn
	mediaandfiles MediaAndFilesClient
	users UsersClient
	authentication AuthenticationClient
	configsync ConfigSyncClient
	contacts ContactsClient
	crypto CryptoClient
	deviceinfo DeviceInfoClient
	eventbus EventBusClient
	obsolete ObsoleteClient
	groups GroupsClient
	integrations IntegrationsClient
	messaging MessagingClient
	privacy PrivacyClient
	profile ProfileClient
	push PushClient
	rawapi RawAPIClient
	registration RegistrationClient
	search SearchClient
	stickers StickersClient
	typingandonline TypingAndOnlineClient
	webrtc WebRTCClient
	spaces SpacesClient
	sequenceandupdates SequenceAndUpdatesClient
}
func NewDialogSDK(cc *grpc.ClientConn) (d *DialogSDK, err error) {
	d = &DialogSDK{}
	d.mediaandfiles = NewMediaAndFilesClient(cc)
	d.users = NewUsersClient(cc)
	d.authentication = NewAuthenticationClient(cc)
	d.configsync = NewConfigSyncClient(cc)
	d.contacts = NewContactsClient(cc)
	d.crypto = NewCryptoClient(cc)
	d.deviceinfo = NewDeviceInfoClient(cc)
	d.eventbus = NewEventBusClient(cc)
	d.obsolete = NewObsoleteClient(cc)
	d.groups = NewGroupsClient(cc)
	d.integrations = NewIntegrationsClient(cc)
	d.messaging = NewMessagingClient(cc)
	d.privacy = NewPrivacyClient(cc)
	d.profile = NewProfileClient(cc)
	d.push = NewPushClient(cc)
	d.rawapi = NewRawAPIClient(cc)
	d.registration = NewRegistrationClient(cc)
	d.search = NewSearchClient(cc)
	d.stickers = NewStickersClient(cc)
	d.typingandonline = NewTypingAndOnlineClient(cc)
	d.webrtc = NewWebRTCClient(cc)
	d.spaces = NewSpacesClient(cc)
	d.sequenceandupdates = NewSequenceAndUpdatesClient(cc)
	d.conn = cc
	err = d.initializeContext()
	return
}
func (d *DialogSDK) MediaAndFilesGetFileUrl(in *RequestGetFileUrl) (*ResponseGetFileUrl, error) {
	return d.mediaandfiles.GetFileUrl(d.internalContext, in)
}
func (d *DialogSDK) MediaAndFilesGetFileUrls(in *RequestGetFileUrls) (*ResponseGetFileUrls, error) {
	return d.mediaandfiles.GetFileUrls(d.internalContext, in)
}
func (d *DialogSDK) MediaAndFilesGetFileUrlBuilder(in *RequestGetFileUrlBuilder) (*ResponseGetFileUrlBuilder, error) {
	return d.mediaandfiles.GetFileUrlBuilder(d.internalContext, in)
}
func (d *DialogSDK) MediaAndFilesGetFileUploadUrl(in *RequestGetFileUploadUrl) (*ResponseGetFileUploadUrl, error) {
	return d.mediaandfiles.GetFileUploadUrl(d.internalContext, in)
}
func (d *DialogSDK) MediaAndFilesCommitFileUpload(in *RequestCommitFileUpload) (*ResponseCommitFileUpload, error) {
	return d.mediaandfiles.CommitFileUpload(d.internalContext, in)
}
func (d *DialogSDK) MediaAndFilesGetFileUploadPartUrl(in *RequestGetFileUploadPartUrl) (*ResponseGetFileUploadPartUrl, error) {
	return d.mediaandfiles.GetFileUploadPartUrl(d.internalContext, in)
}
func (d *DialogSDK) UsersEditUserLocalName(in *RequestEditUserLocalName) (*ResponseSeq, error) {
	return d.users.EditUserLocalName(d.internalContext, in)
}
func (d *DialogSDK) UsersLoadFullUsers(in *RequestLoadFullUsers) (*ResponseLoadFullUsers, error) {
	return d.users.LoadFullUsers(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationStartPhoneAuth(in *RequestStartPhoneAuth) (*ResponseStartPhoneAuth, error) {
	return d.authentication.StartPhoneAuth(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSendCodeByPhoneCall(in *RequestSendCodeByPhoneCall) (*ResponseVoid, error) {
	return d.authentication.SendCodeByPhoneCall(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationStartEmailAuth(in *RequestStartEmailAuth) (*ResponseStartEmailAuth, error) {
	return d.authentication.StartEmailAuth(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationStartAnonymousAuth(in *RequestStartAnonymousAuth) (*ResponseAuth, error) {
	return d.authentication.StartAnonymousAuth(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationStartTokenAuth(in *RequestStartTokenAuth) (*ResponseAuth, error) {
	return d.authentication.StartTokenAuth(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationStartUsernameAuth(in *RequestStartUsernameAuth) (*ResponseStartUsernameAuth, error) {
	return d.authentication.StartUsernameAuth(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationValidateCode(in *RequestValidateCode) (*ResponseAuth, error) {
	return d.authentication.ValidateCode(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationResendCode(in *RequestResendCode) (*ResponseVoid, error) {
	return d.authentication.ResendCode(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationValidatePassword(in *RequestValidatePassword) (*ResponseAuth, error) {
	return d.authentication.ValidatePassword(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationGetOAuth2Params(in *RequestGetOAuth2Params) (*ResponseGetOAuth2Params, error) {
	return d.authentication.GetOAuth2Params(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationCompleteOAuth2(in *RequestCompleteOAuth2) (*ResponseAuth, error) {
	return d.authentication.CompleteOAuth2(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSignUp(in *RequestSignUp) (*ResponseAuth, error) {
	return d.authentication.SignUp(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationGetAuthSessions(in *RequestGetAuthSessions) (*ResponseGetAuthSessions, error) {
	return d.authentication.GetAuthSessions(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationTerminateSession(in *RequestTerminateSession) (*ResponseVoid, error) {
	return d.authentication.TerminateSession(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationTerminateAllSessions(in *RequestTerminateAllSessions) (*ResponseVoid, error) {
	return d.authentication.TerminateAllSessions(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSignOut(in *RequestSignOut) (*ResponseVoid, error) {
	return d.authentication.SignOut(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSignInObsolete(in *RequestSignInObsolete) (*ResponseAuth, error) {
	return d.authentication.SignInObsolete(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSignUpObsolete(in *RequestSignUpObsolete) (*ResponseAuth, error) {
	return d.authentication.SignUpObsolete(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSendAuthCodeObsolete(in *RequestSendAuthCodeObsolete) (*ResponseSendAuthCodeObsolete, error) {
	return d.authentication.SendAuthCodeObsolete(d.internalContext, in)
}
func (d *DialogSDK) AuthenticationSendAuthCallObsolete(in *RequestSendAuthCallObsolete) (*ResponseVoid, error) {
	return d.authentication.SendAuthCallObsolete(d.internalContext, in)
}
func (d *DialogSDK) ConfigSyncGetParameters(in *RequestGetParameters) (*ResponseGetParameters, error) {
	return d.configsync.GetParameters(d.internalContext, in)
}
func (d *DialogSDK) ConfigSyncEditParameter(in *RequestEditParameter) (*ResponseSeq, error) {
	return d.configsync.EditParameter(d.internalContext, in)
}
func (d *DialogSDK) ContactsImportContacts(in *RequestImportContacts) (*ResponseImportContacts, error) {
	return d.contacts.ImportContacts(d.internalContext, in)
}
func (d *DialogSDK) ContactsDeferredImportContacts(in *RequestDeferredImportContacts) (*ResponseDeferredImportContacts, error) {
	return d.contacts.DeferredImportContacts(d.internalContext, in)
}
func (d *DialogSDK) ContactsGetContacts(in *RequestGetContacts) (*ResponseGetContacts, error) {
	return d.contacts.GetContacts(d.internalContext, in)
}
func (d *DialogSDK) ContactsRemoveContact(in *RequestRemoveContact) (*ResponseSeq, error) {
	return d.contacts.RemoveContact(d.internalContext, in)
}
func (d *DialogSDK) ContactsAddContact(in *RequestAddContact) (*ResponseSeq, error) {
	return d.contacts.AddContact(d.internalContext, in)
}
func (d *DialogSDK) ContactsSearchContacts(in *RequestSearchContacts) (*ResponseSearchContacts, error) {
	return d.contacts.SearchContacts(d.internalContext, in)
}
func (d *DialogSDK) CryptoKeyExchange(in *RequestKeyExchange) (*ResponseKeyExchange, error) {
	return d.crypto.KeyExchange(d.internalContext, in)
}
func (d *DialogSDK) DeviceInfoNotifyAboutDeviceInfo(in *RequestNotifyAboutDeviceInfo) (*ResponseVoid, error) {
	return d.deviceinfo.NotifyAboutDeviceInfo(d.internalContext, in)
}
func (d *DialogSDK) EventBusJoinEventBus(in *RequestJoinEventBus) (*ResponseJoinEventBus, error) {
	return d.eventbus.JoinEventBus(d.internalContext, in)
}
func (d *DialogSDK) EventBusKeepAliveEventBus(in *RequestKeepAliveEventBus) (*ResponseVoid, error) {
	return d.eventbus.KeepAliveEventBus(d.internalContext, in)
}
func (d *DialogSDK) EventBusPostToEventBus(in *RequestPostToEventBus) (*ResponseVoid, error) {
	return d.eventbus.PostToEventBus(d.internalContext, in)
}
func (d *DialogSDK) ObsoleteObsolete(in *BytesValue) (*BytesValue, error) {
	return d.obsolete.Obsolete(d.internalContext, in)
}
func (d *DialogSDK) ObsoleteSeqUpdates(in *Empty) (Obsolete_SeqUpdatesClient, error) {
	return d.obsolete.SeqUpdates(d.internalContext, in)
}
func (d *DialogSDK) ObsoleteWeakUpdates() (Obsolete_WeakUpdatesClient, error) {
	return d.obsolete.WeakUpdates(d.internalContext)
}
func (d *DialogSDK) GroupsLoadFullGroups(in *RequestLoadFullGroups) (*ResponseLoadFullGroups, error) {
	return d.groups.LoadFullGroups(d.internalContext, in)
}
func (d *DialogSDK) GroupsLoadMembers(in *RequestLoadMembers) (*ResponseLoadMembers, error) {
	return d.groups.LoadMembers(d.internalContext, in)
}
func (d *DialogSDK) GroupsCreateGroup(in *RequestCreateGroup) (*ResponseCreateGroup, error) {
	return d.groups.CreateGroup(d.internalContext, in)
}
func (d *DialogSDK) GroupsEditGroupTitle(in *RequestEditGroupTitle) (*ResponseSeqDateMid, error) {
	return d.groups.EditGroupTitle(d.internalContext, in)
}
func (d *DialogSDK) GroupsSetGroupShortname(in *RequestSetGroupShortname) (*ResponseSeq, error) {
	return d.groups.SetGroupShortname(d.internalContext, in)
}
func (d *DialogSDK) GroupsEditGroupAvatar(in *RequestEditGroupAvatar) (*ResponseEditGroupAvatar, error) {
	return d.groups.EditGroupAvatar(d.internalContext, in)
}
func (d *DialogSDK) GroupsRemoveGroupAvatar(in *RequestRemoveGroupAvatar) (*ResponseSeqDateMid, error) {
	return d.groups.RemoveGroupAvatar(d.internalContext, in)
}
func (d *DialogSDK) GroupsEditGroupTopic(in *RequestEditGroupTopic) (*ResponseSeqDate, error) {
	return d.groups.EditGroupTopic(d.internalContext, in)
}
func (d *DialogSDK) GroupsEditGroupAbout(in *RequestEditGroupAbout) (*ResponseSeqDate, error) {
	return d.groups.EditGroupAbout(d.internalContext, in)
}
func (d *DialogSDK) GroupsInviteUser(in *RequestInviteUser) (*ResponseSeqDateMid, error) {
	return d.groups.InviteUser(d.internalContext, in)
}
func (d *DialogSDK) GroupsLeaveGroup(in *RequestLeaveGroup) (*ResponseSeqDateMid, error) {
	return d.groups.LeaveGroup(d.internalContext, in)
}
func (d *DialogSDK) GroupsKickUser(in *RequestKickUser) (*ResponseSeqDateMid, error) {
	return d.groups.KickUser(d.internalContext, in)
}
func (d *DialogSDK) GroupsMakeUserAdmin(in *RequestMakeUserAdmin) (*ResponseSeqDate, error) {
	return d.groups.MakeUserAdmin(d.internalContext, in)
}
func (d *DialogSDK) GroupsGetGroupMemberPermissions(in *RequestGetGroupMemberPermissions) (*ResponseGetGroupMemberPermissions, error) {
	return d.groups.GetGroupMemberPermissions(d.internalContext, in)
}
func (d *DialogSDK) GroupsTransferOwnership(in *RequestTransferOwnership) (*ResponseSeqDate, error) {
	return d.groups.TransferOwnership(d.internalContext, in)
}
func (d *DialogSDK) GroupsGetGroupInviteUrl(in *RequestGetGroupInviteUrl) (*ResponseInviteUrl, error) {
	return d.groups.GetGroupInviteUrl(d.internalContext, in)
}
func (d *DialogSDK) GroupsGetGroupInviteUrlBase(in *RequestGetGroupInviteUrlBase) (*ResponseGetGroupInviteUrlBase, error) {
	return d.groups.GetGroupInviteUrlBase(d.internalContext, in)
}
func (d *DialogSDK) GroupsRevokeInviteUrl(in *RequestRevokeInviteUrl) (*ResponseInviteUrl, error) {
	return d.groups.RevokeInviteUrl(d.internalContext, in)
}
func (d *DialogSDK) GroupsJoinGroup(in *RequestJoinGroup) (*ResponseJoinGroup, error) {
	return d.groups.JoinGroup(d.internalContext, in)
}
func (d *DialogSDK) GroupsJoinGroupByPeer(in *RequestJoinGroupByPeer) (*ResponseVoid, error) {
	return d.groups.JoinGroupByPeer(d.internalContext, in)
}
func (d *DialogSDK) GroupsMakeUserAdminObsolete(in *RequestMakeUserAdminObsolete) (*ResponseMakeUserAdminObsolete, error) {
	return d.groups.MakeUserAdminObsolete(d.internalContext, in)
}
func (d *DialogSDK) IntegrationsGetIntegrationToken(in *RequestGetIntegrationToken) (*ResponseIntegrationToken, error) {
	return d.integrations.GetIntegrationToken(d.internalContext, in)
}
func (d *DialogSDK) IntegrationsRevokeIntegrationToken(in *RequestRevokeIntegrationToken) (*ResponseIntegrationToken, error) {
	return d.integrations.RevokeIntegrationToken(d.internalContext, in)
}
func (d *DialogSDK) MessagingDoInteractiveMediaAction(in *RequestDoInteractiveMediaAction) (*ResponseVoid, error) {
	return d.messaging.DoInteractiveMediaAction(d.internalContext, in)
}
func (d *DialogSDK) MessagingSendMessage(in *RequestSendMessage) (*ResponseSeqDate, error) {
	return d.messaging.SendMessage(d.internalContext, in)
}
func (d *DialogSDK) MessagingUpdateMessage(in *RequestUpdateMessage) (*ResponseSeqDate, error) {
	return d.messaging.UpdateMessage(d.internalContext, in)
}
func (d *DialogSDK) MessagingMessageReceived(in *RequestMessageReceived) (*ResponseVoid, error) {
	return d.messaging.MessageReceived(d.internalContext, in)
}
func (d *DialogSDK) MessagingMessageRead(in *RequestMessageRead) (*ResponseVoid, error) {
	return d.messaging.MessageRead(d.internalContext, in)
}
func (d *DialogSDK) MessagingDeleteMessageObsolete(in *RequestDeleteMessageObsolete) (*ResponseSeq, error) {
	return d.messaging.DeleteMessageObsolete(d.internalContext, in)
}
func (d *DialogSDK) MessagingClearChat(in *RequestClearChat) (*ResponseSeq, error) {
	return d.messaging.ClearChat(d.internalContext, in)
}
func (d *DialogSDK) MessagingDeleteChat(in *RequestDeleteChat) (*ResponseSeq, error) {
	return d.messaging.DeleteChat(d.internalContext, in)
}
func (d *DialogSDK) MessagingArchiveChat(in *RequestArchiveChat) (*ResponseSeq, error) {
	return d.messaging.ArchiveChat(d.internalContext, in)
}
func (d *DialogSDK) MessagingMessageSetReaction(in *RequestMessageSetReaction) (*ResponseReactionsResponse, error) {
	return d.messaging.MessageSetReaction(d.internalContext, in)
}
func (d *DialogSDK) MessagingMessageRemoveReaction(in *RequestMessageRemoveReaction) (*ResponseReactionsResponse, error) {
	return d.messaging.MessageRemoveReaction(d.internalContext, in)
}
func (d *DialogSDK) MessagingLoadHistory(in *RequestLoadHistory) (*ResponseLoadHistory, error) {
	return d.messaging.LoadHistory(d.internalContext, in)
}
func (d *DialogSDK) MessagingLoadDialogs(in *RequestLoadDialogs) (*ResponseLoadDialogs, error) {
	return d.messaging.LoadDialogs(d.internalContext, in)
}
func (d *DialogSDK) MessagingFetchDialogIndex(in *RequestFetchDialogIndex) (*ResponseFetchDialogIndex, error) {
	return d.messaging.FetchDialogIndex(d.internalContext, in)
}
func (d *DialogSDK) MessagingLoadArchived(in *RequestLoadArchived) (*ResponseLoadArchived, error) {
	return d.messaging.LoadArchived(d.internalContext, in)
}
func (d *DialogSDK) MessagingLoadGroupedDialogs(in *RequestLoadGroupedDialogs) (*ResponseLoadGroupedDialogs, error) {
	return d.messaging.LoadGroupedDialogs(d.internalContext, in)
}
func (d *DialogSDK) MessagingHideDialog(in *RequestHideDialog) (*ResponseDialogsOrder, error) {
	return d.messaging.HideDialog(d.internalContext, in)
}
func (d *DialogSDK) MessagingShowDialog(in *RequestShowDialog) (*ResponseDialogsOrder, error) {
	return d.messaging.ShowDialog(d.internalContext, in)
}
func (d *DialogSDK) MessagingFavouriteDialog(in *RequestFavouriteDialog) (*ResponseDialogsOrder, error) {
	return d.messaging.FavouriteDialog(d.internalContext, in)
}
func (d *DialogSDK) MessagingUnfavouriteDialog(in *RequestUnfavouriteDialog) (*ResponseDialogsOrder, error) {
	return d.messaging.UnfavouriteDialog(d.internalContext, in)
}
func (d *DialogSDK) MessagingNotifyDialogOpened(in *RequestNotifyDialogOpened) (*ResponseVoid, error) {
	return d.messaging.NotifyDialogOpened(d.internalContext, in)
}
func (d *DialogSDK) MessagingPinMessage(in *RequestPinMessage) (*ResponseSeqDate, error) {
	return d.messaging.PinMessage(d.internalContext, in)
}
func (d *DialogSDK) MessagingUnpinMessage(in *RequestUnpinMessage) (*ResponseSeqDate, error) {
	return d.messaging.UnpinMessage(d.internalContext, in)
}
func (d *DialogSDK) PrivacyBlockUser(in *RequestBlockUser) (*ResponseSeq, error) {
	return d.privacy.BlockUser(d.internalContext, in)
}
func (d *DialogSDK) PrivacyUnblockUser(in *RequestUnblockUser) (*ResponseSeq, error) {
	return d.privacy.UnblockUser(d.internalContext, in)
}
func (d *DialogSDK) PrivacyLoadBlockedUsers(in *RequestLoadBlockedUsers) (*ResponseLoadBlockedUsers, error) {
	return d.privacy.LoadBlockedUsers(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditName(in *RequestEditName) (*ResponseSeq, error) {
	return d.profile.EditName(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditNickName(in *RequestEditNickName) (*ResponseSeq, error) {
	return d.profile.EditNickName(d.internalContext, in)
}
func (d *DialogSDK) ProfileCheckNickName(in *RequestCheckNickName) (*ResponseBool, error) {
	return d.profile.CheckNickName(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditAbout(in *RequestEditAbout) (*ResponseSeq, error) {
	return d.profile.EditAbout(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditAvatar(in *RequestEditAvatar) (*ResponseEditAvatar, error) {
	return d.profile.EditAvatar(d.internalContext, in)
}
func (d *DialogSDK) ProfileRemoveAvatar(in *RequestRemoveAvatar) (*ResponseSeq, error) {
	return d.profile.RemoveAvatar(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditMyTimeZone(in *RequestEditMyTimeZone) (*ResponseSeq, error) {
	return d.profile.EditMyTimeZone(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditMyPreferredLanguages(in *RequestEditMyPreferredLanguages) (*ResponseSeq, error) {
	return d.profile.EditMyPreferredLanguages(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditSex(in *RequestEditSex) (*ResponseSeq, error) {
	return d.profile.EditSex(d.internalContext, in)
}
func (d *DialogSDK) ProfileEditCustomProfile(in *RequestEditCustomProfile) (*ResponseSeq, error) {
	return d.profile.EditCustomProfile(d.internalContext, in)
}
func (d *DialogSDK) ProfileChangeUserStatus(in *RequestChangeUserStatus) (*ResponseSeq, error) {
	return d.profile.ChangeUserStatus(d.internalContext, in)
}
func (d *DialogSDK) PushRegisterGooglePush(in *RequestRegisterGooglePush) (*ResponseVoid, error) {
	return d.push.RegisterGooglePush(d.internalContext, in)
}
func (d *DialogSDK) PushUnregisterGooglePush(in *RequestUnregisterGooglePush) (*ResponseVoid, error) {
	return d.push.UnregisterGooglePush(d.internalContext, in)
}
func (d *DialogSDK) PushRegisterApplePush(in *RequestRegisterApplePush) (*ResponseVoid, error) {
	return d.push.RegisterApplePush(d.internalContext, in)
}
func (d *DialogSDK) PushUnregisterApplePush(in *RequestUnregisterApplePush) (*ResponseVoid, error) {
	return d.push.UnregisterApplePush(d.internalContext, in)
}
func (d *DialogSDK) PushRegisterApplePushKit(in *RequestRegisterApplePushKit) (*ResponseVoid, error) {
	return d.push.RegisterApplePushKit(d.internalContext, in)
}
func (d *DialogSDK) PushUnregisterApplePushKit(in *RequestUnregisterApplePushKit) (*ResponseVoid, error) {
	return d.push.UnregisterApplePushKit(d.internalContext, in)
}
func (d *DialogSDK) PushRegisterApplePushToken(in *RequestRegisterApplePushToken) (*ResponseVoid, error) {
	return d.push.RegisterApplePushToken(d.internalContext, in)
}
func (d *DialogSDK) PushUnregisterApplePushToken(in *RequestUnregisterApplePushToken) (*ResponseVoid, error) {
	return d.push.UnregisterApplePushToken(d.internalContext, in)
}
func (d *DialogSDK) RawAPIRawRequest(in *RequestRawRequest) (*ResponseRawRequest, error) {
	return d.rawapi.RawRequest(d.internalContext, in)
}
func (d *DialogSDK) RegistrationExchangeAuthIdForToken(in *RequestExchangeAuthIdForToken) (*ResponseDeviceRequest, error) {
	return d.registration.ExchangeAuthIdForToken(d.internalContext, in)
}
func (d *DialogSDK) RegistrationRegisterDevice(in *RequestRegisterDevice) (*ResponseDeviceRequest, error) {
	return d.registration.RegisterDevice(d.internalContext, in)
}
func (d *DialogSDK) RegistrationRegisterDeprecatedDevice(in *RegisterDeprecatedDeviceRequest) (*ResponseDeviceRequest, error) {
	return d.registration.RegisterDeprecatedDevice(d.internalContext, in)
}
func (d *DialogSDK) SearchPeerSearch(in *RequestPeerSearch) (*ResponsePeerSearch, error) {
	return d.search.PeerSearch(d.internalContext, in)
}
func (d *DialogSDK) SearchResolvePeer(in *RequestResolvePeer) (*ResponseResolvePeer, error) {
	return d.search.ResolvePeer(d.internalContext, in)
}
func (d *DialogSDK) SearchMessageSearch(in *RequestMessageSearch) (*ResponseMessageSearchResponse, error) {
	return d.search.MessageSearch(d.internalContext, in)
}
func (d *DialogSDK) SearchMessageSearchMore(in *RequestMessageSearchMore) (*ResponseMessageSearchResponse, error) {
	return d.search.MessageSearchMore(d.internalContext, in)
}
func (d *DialogSDK) SearchSimpleSearch(in *RequestSimpleSearch) (*ResponseMessageSearchResponse, error) {
	return d.search.SimpleSearch(d.internalContext, in)
}
func (d *DialogSDK) SearchSimpleSearchMore(in *RequestSimpleSearchMore) (*ResponseMessageSearchResponse, error) {
	return d.search.SimpleSearchMore(d.internalContext, in)
}
func (d *DialogSDK) SearchAutocompleteSuggestions(in *RequestFieldAutocomplete) (*ResponseFieldAutocomplete, error) {
	return d.search.AutocompleteSuggestions(d.internalContext, in)
}
func (d *DialogSDK) SearchLoadUserSearchByPredicatesResults(in *RequestLoadUserSearchByPredicatesResults) (*ResponseLoadUserSearchByPredicatesResults, error) {
	return d.search.LoadUserSearchByPredicatesResults(d.internalContext, in)
}
func (d *DialogSDK) SearchLoadUserSearchByPredicatesCount(in *RequestLoadUserSearchByPredicatesCount) (*ResponseLoadUserSearchByPredicatesCount, error) {
	return d.search.LoadUserSearchByPredicatesCount(d.internalContext, in)
}
func (d *DialogSDK) StickersLoadOwnStickers(in *RequestLoadOwnStickers) (*ResponseLoadOwnStickers, error) {
	return d.stickers.LoadOwnStickers(d.internalContext, in)
}
func (d *DialogSDK) StickersLoadAcesssibleStickers(in *RequestLoadAcesssibleStickers) (*ResponseLoadAcesssibleStickers, error) {
	return d.stickers.LoadAcesssibleStickers(d.internalContext, in)
}
func (d *DialogSDK) StickersAddStickerPackReference(in *RequestAddStickerPackReference) (*ResponseSeq, error) {
	return d.stickers.AddStickerPackReference(d.internalContext, in)
}
func (d *DialogSDK) StickersRemoveStickerPackReference(in *RequestRemoveStickerPackReference) (*ResponseSeq, error) {
	return d.stickers.RemoveStickerPackReference(d.internalContext, in)
}
func (d *DialogSDK) StickersAddStickerCollection(in *RequestAddStickerCollection) (*ResponseSeq, error) {
	return d.stickers.AddStickerCollection(d.internalContext, in)
}
func (d *DialogSDK) StickersRemoveStickerCollection(in *RequestRemoveStickerCollection) (*ResponseSeq, error) {
	return d.stickers.RemoveStickerCollection(d.internalContext, in)
}
func (d *DialogSDK) StickersLoadStickerCollection(in *RequestLoadStickerCollection) (*ResponseLoadStickerCollection, error) {
	return d.stickers.LoadStickerCollection(d.internalContext, in)
}
func (d *DialogSDK) TypingAndOnlineTyping(in *RequestTyping) (*ResponseVoid, error) {
	return d.typingandonline.Typing(d.internalContext, in)
}
func (d *DialogSDK) TypingAndOnlineStopTyping(in *RequestStopTyping) (*ResponseVoid, error) {
	return d.typingandonline.StopTyping(d.internalContext, in)
}
func (d *DialogSDK) TypingAndOnlineSetOnline(in *RequestSetOnline) (*ResponseVoid, error) {
	return d.typingandonline.SetOnline(d.internalContext, in)
}
func (d *DialogSDK) TypingAndOnlinePauseNotifications(in *RequestPauseNotifications) (*ResponseVoid, error) {
	return d.typingandonline.PauseNotifications(d.internalContext, in)
}
func (d *DialogSDK) TypingAndOnlineRestoreNotifications(in *RequestRestoreNotifications) (*ResponseVoid, error) {
	return d.typingandonline.RestoreNotifications(d.internalContext, in)
}
func (d *DialogSDK) WebRTCGetCallInfo(in *RequestGetCallInfo) (*ResponseGetCallInfo, error) {
	return d.webrtc.GetCallInfo(d.internalContext, in)
}
func (d *DialogSDK) WebRTCLoadCalls(in *RequestLoadCalls) (*ResponseLoadCalls, error) {
	return d.webrtc.LoadCalls(d.internalContext, in)
}
func (d *DialogSDK) WebRTCDoCall(in *RequestDoCall) (*ResponseDoCall, error) {
	return d.webrtc.DoCall(d.internalContext, in)
}
func (d *DialogSDK) WebRTCJoinCall(in *RequestJoinCall) (*ResponseVoid, error) {
	return d.webrtc.JoinCall(d.internalContext, in)
}
func (d *DialogSDK) WebRTCRejectCall(in *RequestRejectCall) (*ResponseVoid, error) {
	return d.webrtc.RejectCall(d.internalContext, in)
}
func (d *DialogSDK) WebRTCChangeCallDisplayName(in *RequestChangeCallDisplayName) (*ResponseVoid, error) {
	return d.webrtc.ChangeCallDisplayName(d.internalContext, in)
}
func (d *DialogSDK) WebRTCDeleteCall(in *RequestDeleteCall) (*ResponseVoid, error) {
	return d.webrtc.DeleteCall(d.internalContext, in)
}
func (d *DialogSDK) SpacesCreateSpace(in *RequestCreateSpace) (*ResponseSpace, error) {
	return d.spaces.CreateSpace(d.internalContext, in)
}
func (d *DialogSDK) SpacesDeleteSpace(in *RequestDeleteSpace) (*ResponseSpace, error) {
	return d.spaces.DeleteSpace(d.internalContext, in)
}
func (d *DialogSDK) SpacesSetTitle(in *RequestSetTitle) (*ResponseSpace, error) {
	return d.spaces.SetTitle(d.internalContext, in)
}
func (d *DialogSDK) SpacesSetShortname(in *RequestSetShortname) (*ResponseSpace, error) {
	return d.spaces.SetShortname(d.internalContext, in)
}
func (d *DialogSDK) SpacesSetAbout(in *RequestSetAbout) (*ResponseSpace, error) {
	return d.spaces.SetAbout(d.internalContext, in)
}
func (d *DialogSDK) SpacesSetAvatar(in *RequestSetAvatar) (*ResponseSpace, error) {
	return d.spaces.SetAvatar(d.internalContext, in)
}
func (d *DialogSDK) SpacesLoadSpaces(in *RequestLoadSpaces) (*ResponseLoadSpaces, error) {
	return d.spaces.LoadSpaces(d.internalContext, in)
}
func (d *DialogSDK) SpacesLoadSpaceMembers(in *RequestStreamSpaceMembers) (Spaces_LoadSpaceMembersClient, error) {
	return d.spaces.LoadSpaceMembers(d.internalContext, in)
}
func (d *DialogSDK) SpacesInvite(in *RequestSpaceInvite) (*ResponseSpaceMember, error) {
	return d.spaces.Invite(d.internalContext, in)
}
func (d *DialogSDK) SpacesKick(in *RequestSpaceKick) (*ResponseSpaceMember, error) {
	return d.spaces.Kick(d.internalContext, in)
}
func (d *DialogSDK) SpacesLeave(in *RequestSpaceLeave) (*ResponseSpaceMember, error) {
	return d.spaces.Leave(d.internalContext, in)
}
func (d *DialogSDK) SpacesGetSpaceInviteUrl(in *RequestGetSpaceInviteUrl) (*ResponseSpaceInviteUrl, error) {
	return d.spaces.GetSpaceInviteUrl(d.internalContext, in)
}
func (d *DialogSDK) SpacesRevokeSpaceInviteUrl(in *RequestRevokeSpaceInviteUrl) (*ResponseSpaceInviteUrl, error) {
	return d.spaces.RevokeSpaceInviteUrl(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesGetState(in *RequestGetState) (*ResponseSeq, error) {
	return d.sequenceandupdates.GetState(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesGetDifference(in *RequestGetDifference) (*ResponseGetDifference, error) {
	return d.sequenceandupdates.GetDifference(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesGetDialogsDifference(in *RequestGetDialogsDifference) (*ResponseGetDialogsDifference, error) {
	return d.sequenceandupdates.GetDialogsDifference(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesGetReferencedEntitites(in *RequestGetReferencedEntitites) (*ResponseGetReferencedEntitites, error) {
	return d.sequenceandupdates.GetReferencedEntitites(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesSubscribeToOnline(in *RequestSubscribeToOnline) (*ResponseVoid, error) {
	return d.sequenceandupdates.SubscribeToOnline(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesSubscribeFromOnline(in *RequestSubscribeFromOnline) (*ResponseVoid, error) {
	return d.sequenceandupdates.SubscribeFromOnline(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesSubscribeToGroupOnline(in *RequestSubscribeToGroupOnline) (*ResponseVoid, error) {
	return d.sequenceandupdates.SubscribeToGroupOnline(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesSubscribeFromGroupOnline(in *RequestSubscribeFromGroupOnline) (*ResponseVoid, error) {
	return d.sequenceandupdates.SubscribeFromGroupOnline(d.internalContext, in)
}
func (d *DialogSDK) SequenceAndUpdatesSeqUpdates(in *Empty) (SequenceAndUpdates_SeqUpdatesClient, error) {
	return d.sequenceandupdates.SeqUpdates(d.internalContext, in)
}
