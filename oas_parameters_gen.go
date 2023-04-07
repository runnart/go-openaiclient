// Code generated by ogen, DO NOT EDIT.

package go_openaiclient

import (
	"net/http"
	"net/url"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
	"github.com/ogen-go/ogen/validate"
)

// CancelFineTuneParams is parameters of cancelFineTune operation.
type CancelFineTuneParams struct {
	// The ID of the fine-tune job to cancel.
	FineTuneID string
}

func unpackCancelFineTuneParams(packed middleware.Parameters) (params CancelFineTuneParams) {
	{
		key := middleware.ParameterKey{
			Name: "fine_tune_id",
			In:   "path",
		}
		params.FineTuneID = packed[key].(string)
	}
	return params
}

func decodeCancelFineTuneParams(args [1]string, argsEscaped bool, r *http.Request) (params CancelFineTuneParams, _ error) {
	// Decode path: fine_tune_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "fine_tune_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FineTuneID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "fine_tune_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteFileParams is parameters of deleteFile operation.
type DeleteFileParams struct {
	// The ID of the file to use for this request.
	FileID string
}

func unpackDeleteFileParams(packed middleware.Parameters) (params DeleteFileParams) {
	{
		key := middleware.ParameterKey{
			Name: "file_id",
			In:   "path",
		}
		params.FileID = packed[key].(string)
	}
	return params
}

func decodeDeleteFileParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteFileParams, _ error) {
	// Decode path: file_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "file_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FileID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "file_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DeleteModelParams is parameters of deleteModel operation.
type DeleteModelParams struct {
	// The model to delete.
	Model string
}

func unpackDeleteModelParams(packed middleware.Parameters) (params DeleteModelParams) {
	{
		key := middleware.ParameterKey{
			Name: "model",
			In:   "path",
		}
		params.Model = packed[key].(string)
	}
	return params
}

func decodeDeleteModelParams(args [1]string, argsEscaped bool, r *http.Request) (params DeleteModelParams, _ error) {
	// Decode path: model.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "model",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Model = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "model",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// DownloadFileParams is parameters of downloadFile operation.
type DownloadFileParams struct {
	// The ID of the file to use for this request.
	FileID string
}

func unpackDownloadFileParams(packed middleware.Parameters) (params DownloadFileParams) {
	{
		key := middleware.ParameterKey{
			Name: "file_id",
			In:   "path",
		}
		params.FileID = packed[key].(string)
	}
	return params
}

func decodeDownloadFileParams(args [1]string, argsEscaped bool, r *http.Request) (params DownloadFileParams, _ error) {
	// Decode path: file_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "file_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FileID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "file_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// ListFineTuneEventsParams is parameters of listFineTuneEvents operation.
type ListFineTuneEventsParams struct {
	// The ID of the fine-tune job to get events for.
	FineTuneID string
	// Whether to stream events for the fine-tune job. If set to true,
	// events will be sent as data-only
	// [server-sent events](https://developer.mozilla.
	// org/en-US/docs/Web/API/Server-sent_events/Using_server-sent_events#Event_stream_format)
	// as they become available. The stream will terminate with a
	// `data: [DONE]` message when the job is finished (succeeded, cancelled,
	// or failed).
	// If set to false, only events generated so far will be returned.
	Stream OptBool
}

func unpackListFineTuneEventsParams(packed middleware.Parameters) (params ListFineTuneEventsParams) {
	{
		key := middleware.ParameterKey{
			Name: "fine_tune_id",
			In:   "path",
		}
		params.FineTuneID = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "stream",
			In:   "query",
		}
		if v, ok := packed[key]; ok {
			params.Stream = v.(OptBool)
		}
	}
	return params
}

func decodeListFineTuneEventsParams(args [1]string, argsEscaped bool, r *http.Request) (params ListFineTuneEventsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode path: fine_tune_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "fine_tune_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FineTuneID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "fine_tune_id",
			In:   "path",
			Err:  err,
		}
	}
	// Set default value for query: stream.
	{
		val := bool(false)
		params.Stream.SetTo(val)
	}
	// Decode query: stream.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "stream",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotStreamVal bool
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToBool(val)
					if err != nil {
						return err
					}

					paramsDotStreamVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Stream.SetTo(paramsDotStreamVal)
				return nil
			}); err != nil {
				return err
			}
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "stream",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RetrieveFileParams is parameters of retrieveFile operation.
type RetrieveFileParams struct {
	// The ID of the file to use for this request.
	FileID string
}

func unpackRetrieveFileParams(packed middleware.Parameters) (params RetrieveFileParams) {
	{
		key := middleware.ParameterKey{
			Name: "file_id",
			In:   "path",
		}
		params.FileID = packed[key].(string)
	}
	return params
}

func decodeRetrieveFileParams(args [1]string, argsEscaped bool, r *http.Request) (params RetrieveFileParams, _ error) {
	// Decode path: file_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "file_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FileID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "file_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RetrieveFineTuneParams is parameters of retrieveFineTune operation.
type RetrieveFineTuneParams struct {
	// The ID of the fine-tune job.
	FineTuneID string
}

func unpackRetrieveFineTuneParams(packed middleware.Parameters) (params RetrieveFineTuneParams) {
	{
		key := middleware.ParameterKey{
			Name: "fine_tune_id",
			In:   "path",
		}
		params.FineTuneID = packed[key].(string)
	}
	return params
}

func decodeRetrieveFineTuneParams(args [1]string, argsEscaped bool, r *http.Request) (params RetrieveFineTuneParams, _ error) {
	// Decode path: fine_tune_id.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "fine_tune_id",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.FineTuneID = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "fine_tune_id",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}

// RetrieveModelParams is parameters of retrieveModel operation.
type RetrieveModelParams struct {
	// The ID of the model to use for this request.
	Model string
}

func unpackRetrieveModelParams(packed middleware.Parameters) (params RetrieveModelParams) {
	{
		key := middleware.ParameterKey{
			Name: "model",
			In:   "path",
		}
		params.Model = packed[key].(string)
	}
	return params
}

func decodeRetrieveModelParams(args [1]string, argsEscaped bool, r *http.Request) (params RetrieveModelParams, _ error) {
	// Decode path: model.
	if err := func() error {
		param := args[0]
		if argsEscaped {
			unescaped, err := url.PathUnescape(args[0])
			if err != nil {
				return errors.Wrap(err, "unescape path")
			}
			param = unescaped
		}
		if len(param) > 0 {
			d := uri.NewPathDecoder(uri.PathDecoderConfig{
				Param:   "model",
				Value:   param,
				Style:   uri.PathStyleSimple,
				Explode: false,
			})

			if err := func() error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Model = c
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return validate.ErrFieldRequired
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "model",
			In:   "path",
			Err:  err,
		}
	}
	return params, nil
}
